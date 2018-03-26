package gopool

import (
	"time"
	"sync"
	"fmt"
)

//TODO 最好能够根据可用内存大小实现动态更改
const DefaultGoroutineSize = 256 * 1024 / 8  //(8KB/goroutine)
const DefaultGoroutineExpire = 10 * time.Second

var ErrLack = fmt.Errorf("池中没有可使用的goroutine.")


/* GR goroutine 池*/
type GRpool struct {
	size 			int32   		 // 限制最大的goroutine数量/协程数/worker数量
	expire 			time.Duration    // goroutine过期时间(秒)
	grCount         int32   		 // 当前goroutine数量(非任务数)
	stopCh          chan struct{}    // 池关闭事件(用于池相关异步协程通知)
	lock            sync.Mutex       // 互斥锁
	ready           []*goroutineChan // 空闲的 goroutines
	mustStop        bool			 // 关闭 goroutine 池的标记
	goroutineChanPool sync.Pool
}

type goroutineChan struct {
	lastUseTime time.Time
	ch 			chan func()
}


func NewGRpool(size int32, expire time.Duration) *GRpool{
	p  := new(GRpool)

	if size <= 0 {
		p.size = DefaultGoroutineSize
	} else {
		p.size = size
	}

	if expire <= 0 {
		p.expire = DefaultGoroutineExpire
	} else {
		p.expire = expire
	}

	p.start()
	return p
}

func (p *GRpool) start() {
	if p.stopCh != nil {
		panic("BUG: goroutine池已经启动了")
	}

	p.stopCh = make(chan struct{})
	go func() {
		var scratch []*goroutineChan
		for {
			p.clean(&scratch)
			select {
			case <- p.stopCh: //TODO goutil > stopCh
				return
			default:
				time.Sleep(p.expire)
			}
		}
	}()
}

/*
停止GRpool。
如果在调用'Stop'后调用'Go'，将不再重用goroutine。
*/
func (p *GRpool) Stop() {
	if p.stopCh == nil {
		panic("BUG:GRpool 没有启动")
	}
	close(p.stopCh)      //关闭 chan
	p.stopCh = nil

	//停止空闲的goroutine
	p.lock.Lock()
	// 这里
	ready := p.ready
	for i,ch := range ready {
		ch.ch <- nil
		ready[i] = nil
	}

	//TODO p.ready = []*goroutineChan{} 这样不行?
	p.ready = ready[:0]
	p.mustStop = true
	p.lock.Unlock()
}

// 清理过期的 goroutines
func (p *GRpool) clean (scratch *[]*goroutineChan) {
	currentTime := time.Now()

	p.lock.Lock()
	ready := p.ready
	n := len(p.ready)
	i := 0
	for i < n && currentTime.Sub(ready[i].lastUseTime) > p.expire {
		i++
	}
	*scratch = append((*scratch)[:0], ready[:i]...)
	if i > 0 {
		m := copy(ready, ready[i:])
		for i = m; i < n; i++ {
			ready[i] = nil
		}
		p.ready = ready[:m]
	}

	p.lock.Unlock()

	// Notify obsolete goroutines to stop.
	// This notification must be outside the gp.lock, since ch.ch
	// may be blocking and may consume a lot of time if many goroutines
	// are located on non-local CPUs.
	tmp := *scratch
	for i, ch := range tmp {
		ch.ch <- nil
		tmp[i] = nil
	}
}


func (p *GRpool) Go(fn func()) error {
	ch := p.getCh()
	if ch == nil {
		return ErrLack
	}
	ch.ch <- fn
	return nil
}

func(p *GRpool) getCh() *goroutineChan {
	var ch *goroutineChan
	createGoroutine := false

	p.lock.Lock()
	n := len(p.ready) - 1
	if n < 0 {
		if p.grCount < p.size {
			createGoroutine = true
			p.grCount++
		}
	} else {
			ch = p.ready[n]
			p.ready[n] = nil
			p.ready = p.ready[:n]

	}
	p.lock.Unlock()

	if ch == nil {
		if !createGoroutine {
			return nil
		}
		vch := p.goroutineChanPool.Get()
		if vch == nil {
			vch = &goroutineChan{
				ch: make(chan func(),1),
			}
		}
		ch = vch.(*goroutineChan)
		go func() {

			p.goroutineChanPool.Put(vch)
		}()
	}

	return ch
}


func (p *GRpool) goroutineFunc(ch *goroutineChan) {
	//TODO 这里为啥要用 for
	for fn := range ch.ch {
		if fn == nil {
			break
		}
		fn()
		if !p.release(ch) {
			break
		}
	}

	p.lock.Lock()
	p.grCount--
	p.lock.Unlock()
}


func (p *GRpool) release(ch *goroutineChan) bool{
	ch.lastUseTime = time.Now()
	p.lock.Lock()
	if p.mustStop {
		p.lock.Unlock()
		return false
	}
	p.ready = append(p.ready,ch)
	p.lock.Unlock()
	return true
}