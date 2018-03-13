package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10) //缓存个数为10
	for i := 0; i < 10; i++ {
		go Go(c, i)
	} //执行了10个线程

	for i := 0; i < 10; i++ {
		<-c
	} //在这里取10次
}
func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true //在这里存一次
}