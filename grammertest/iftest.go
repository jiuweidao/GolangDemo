package main

import "fmt"

/*
1.if语句里面允许再if里面声明单一变量
 */

func main() {
	if b:=3; 8> b {
		fmt.Println("a 比 10 大")
	} else {
		fmt.Println("a 比 10 小")
	}
}
