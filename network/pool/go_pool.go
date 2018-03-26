package pool

import (
	"time"
	"sync"
)

type GoPool struct {
	maxGoroutinesAmount      int     //goroutine 最大数量
	maxGoroutineIdleDuration time.Duration   //goroutine 最大空闲时间(猜测)

	lock 				sync.Mutex   //互斥锁
	goroutinesCount 	int          //当前的
	mustStop 			bool
	ready 				[]*goroutingChan
	stopCh 				chan struct{}
	goroutingChanPool 	sync.Pool
}


type goroutingChan struct {
	lastUseTime time.Time
	ch chan func()
}