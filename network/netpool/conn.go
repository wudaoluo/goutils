package netpool

import "sync"


// ConnPool 存放链接信息
type ConnPool struct {
	mu sync.Mutex
}