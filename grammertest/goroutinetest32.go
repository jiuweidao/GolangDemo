package main

import (
	_ "fmt"
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 1;x<4; x++ {
			naturals <- x
			if x == 3{
				close(naturals)
				/*
					当一个channel被关闭后，再向该channel发送数据将导致panic异常。
				当一个被关闭的channel中已经发送的数据都被成功接收后，
				后续的接收操作将不再阻塞，它们会立即返回一个零值。
				关闭上面例子中的naturals变量对应的channel并不能终止循环，
				它依然会收到一个永无休止的零值序列，然后将它们发送给打印者goroutine。
				 */
			}
		}
	}()

	// Squarer
	go func() {
		for {
			x,ok := <-naturals
			if ok {
				squares <- x * x
			}else {
				close(squares)
				break
			}


		}
	}()

	// Printer (in main goroutine)
	for {
		//var i int

		if i,ok := <-squares;ok{
			fmt.Println(i,ok)
		}else {
			fmt.Println(i,ok)
			break
		}

	}
}