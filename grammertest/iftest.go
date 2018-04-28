package main

import "fmt"

/*
1.if语句里面允许再if里面声明单一变量
 */

func main() {

	a:=3
	if a==2 {
		fmt.Println("a ==2")
	} else if a==3{
		fmt.Println("a =3")
	}
}
