package main

import (
	"time"
	"fmt"
)

/*
	https://tonybai.com/2016/12/21/how-to-use-timer-reset-in-golang-correctly
*/

//
var c chan string
//
//func main() {
//	select {
//	case  <-c:
//		fmt.Println("aaa")
//	case <-time.After(5 * time.Second):
//		fmt.Println("timed out")
//	}
//
//	time.Sleep(10*time.Second)
//
//}


func main() {
	t := time.NewTimer(1 * time.Second)
	for{

		select {
		case <-c:
			fmt.Println("aaa")
		case <-t.C:
			fmt.Println("timed out")
		}
		t.Reset(1 * time.Second)
	}

}
