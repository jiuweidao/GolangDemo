package main

import "fmt"

var quit chan int // 只开一个信道

func foo(id int) {
	fmt.Println("读入",id)
	quit <- id // ok, finished
}

func main() {
	count := 10
	quit = make(chan int) // 无缓冲

	for i := 0; i < count; i++ {
		go foo(i)
	}

	for i := 0; i < count; i++ {
		fmt.Println("读出",i)
		<- quit
	}
	/*count := 10
	quit = make(chan int,10) // 无缓冲

	for i := 0; i < count; i++ {
		go foo(i)
	}
	for i := 0; i < 1; i++ {
		<- quit
	}*/

}
