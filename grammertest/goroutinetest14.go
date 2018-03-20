package main

import "fmt"

func main() {
	/*
		如果你执行了上面的代码，会报死锁错误的，
	原因是range不等到信道关闭是不会结束读取的。
	也就是如果 缓冲信道干涸了，
	那么range就会阻塞当前goroutine, 所以死锁咯
	close(ch)
	被关闭的信道会禁止数据流入, 是只读的。
	我们仍然可以从关闭的信道中取出数据，但是不能再写入数据了。
	*/
	/*ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	//close(ch)

	for v := range ch {
		fmt.Println(v)
	}*/







	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for v := range ch {
		fmt.Println(v)
		/*if len(ch) <= 0 { // 如果现有数据量为0，跳出循环
			break
		}*/
	}
}
