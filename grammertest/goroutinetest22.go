package main

import (
	"time"
	"fmt"
)

func saySomething(str string) {
	for  {
		time.Sleep(time.Millisecond * 3)
		fmt.Println(str)
	}
}

func main() {
	// 启动一个goroutine线程
	go saySomething("Hello")
	//saySomething("World")

	time.Sleep(time.Millisecond * 12)
}
