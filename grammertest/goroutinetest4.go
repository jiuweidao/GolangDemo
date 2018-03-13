package main

import (
	"fmt"
)

func main() {
	c := make(chan bool, 1) //这就成了有缓存的channel了
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
	}()
	<-c
}