package glist

import (
	"container/list"
	"sync"
)


//变长双向链表
type List struct {
	list *list.List
	mu sync.RWMutex
}


//获取一个变长链表指针
func New() *List {
	return &List{list: list.New()}
}


