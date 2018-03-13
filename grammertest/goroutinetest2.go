package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	//make(chan [所存储的数据类型])，直接make的都是双向通道，又可以存又可以取
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true //把true给到c
		close(c)
	}()

	for v := range c {

		fmt.Println(v)
	}//一定要注意在某处把channel关闭掉，否则会成为死锁，无限等待
}

//这个函数的流程是这样的：
// go关键字并发了一个匿名函数，
// 接着执行for-range，
// 此时for-range的执行需要等待一个c，在匿名函数中，
// true给到c之后，立刻for-range可以执行了，
// 输出了这个v（v就是c），
// 并且立刻又进入等待下一个c的状态，
// 在匿名函数中，此时关闭了c，for-range立刻接收到了c被关闭的信号，整个函数结束。

