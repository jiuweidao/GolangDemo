package main

import (
	"fmt"
//	"time"
//	"time"
	"time"
)

func getInfo(name string, age int) {
	/*go func(){
		fmt.Printf("name: %s,age: %d \n", name, age)
	}*/
	go func() {
		fmt.Printf("name: %s,age: %d \n", name, age)
	}()
}
func main() {
	getInfo("li", 12)
	//go getInfo("sun", 45)
	fmt.Println("start")
	time.Sleep(1000 * 1000) //
	// 若不等待主线程就结束了
}

