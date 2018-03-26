package netpool

import "fmt"

var ErrClose = fmt.Errorf("连接池已经关闭")


type Pool interface {
	// 从池中获取一个新的连接,池中没有闲置连接,将创建一个新连接
	Get() (interface{}, error)

	//
	Put(interface{}) error

	// 关闭连接池和连接池中的所有连接
	Close()

	// 返回连接池中 连接的数量
	Len() int

	// Release 释放连接池中所有链接
	Release()

}