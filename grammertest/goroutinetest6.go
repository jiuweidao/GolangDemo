package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置使用的CPU核数（获取本台计算机的CPU核数）
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go1(c, i)
	}
	//当一个核的时候还是按部就班的一个线程一个线程开始执行，
	// 可是现在变成多核以后，指不定先执行哪个后执行哪个，可能先执行i=5的，也可能先执行i=8的，
	// 即分配任务不定的特性
	<-c
}
func Go1(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	if index == 9 {
		c <- true //因此这个地方就错误了，有时候index=9并不是最后一个执行的。
	}
}