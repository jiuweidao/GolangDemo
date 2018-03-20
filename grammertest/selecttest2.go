package main

import (
	"time"
	"fmt"
)
/*
没有defalt的，
select语句会一直等待，
直到某个case里的IO操作可以进行,
如果由case在其他goroutine有需要的数据，在传递数据则，会一致等待
如果有defalt则不会等待
 */
func f1(ch chan int) {
	time.Sleep(time.Second * 5)
	ch <- 1
}
func f2(ch chan int) {
	time.Sleep(time.Second * 10)
	ch <- 1
}
func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go f1(ch1)
	go f2(ch2)
	select {
	case a:=<-ch1:
		fmt.Println("The first case is selected.",a)
	case b:=<-ch2:
		fmt.Println("The second case is selected.",b)

	default:
		fmt.Println("i dont want to wait")
	}
}
