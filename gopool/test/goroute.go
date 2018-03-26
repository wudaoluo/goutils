package test


/*这个只不过限制了的 goroutine数量, goroutine并没有复用*/

import (
	"fmt"
	"time"
)

type Pool struct {
	work chan func()
	sem chan struct{}
}


func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *Pool) Schedule(task func()) error {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
	return nil
}

func (p *Pool) worker(task func()) {
	defer func() { <-p.sem }()
	for {
		task()
		task = <-p.work
	}
}

func main() {
	pool := New(5)
	for i := 0; i < 10; i++ {
		i:=i
		pool.Schedule(func() {
			fmt.Println("i:",i)
			time.Sleep(3*time.Second)
		})
	}

	time.Sleep(time.Second*3)
}
