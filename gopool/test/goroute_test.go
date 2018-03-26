package test

import (
	"sync"
	"fmt"
	"runtime"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int) // 断言
	p.Put(1)
	runtime.GC()  //gc 回收了
	b := p.Get().(int)
	c := p.Get().(int)
	fmt.Println(a,b,c)
}
