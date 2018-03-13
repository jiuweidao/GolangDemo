package main

import (
	"fmt"
)
/*
无缓存，阻塞
 */
func main() {
	c := make(chan bool) //这里为无缓存的channel
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
	}()
	<-c
	//运行到这里的时候该取了，
	// 可实际还没有放进去，所以在等待，
	// 等待放入之后再执行取的操作，
	// 无缓存是阻塞的，c里的东西不被读掉，程序无法往下执行
}