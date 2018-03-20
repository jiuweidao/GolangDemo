package main

import (
	"fmt"
	"time"
)

var quit chan int // 只开一个信道

func foo(id int) {
	fmt.Println(id,"开始goroutine")
	quit <- id // ok, finished
	fmt.Println("读入",id)
}

func main() {
	count := 10
	quit = make(chan int) // 无缓冲

	for i := 0; i < count; i++ {
		go foo(i)
	}
	fmt.Println("休眠5-",)
	time.Sleep(5* time.Second)
	fmt.Println("读出",<- quit)
	fmt.Println("休眠3-",)
	time.Sleep(3* time.Second)

	/*for i := 0; i < count; i++ {
		fmt.Println("读出",i)
		<- quit
	}*/
	/*count := 10
	quit = make(chan int,10) // 无缓冲

	for i := 0; i < count; i++ {
		go foo(i)
	}
	for i := 0; i < 1; i++ {
		<- quit
	}*/

}
