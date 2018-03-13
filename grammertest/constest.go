package main

import (
	"fmt"
	"unsafe"
)

const (
	i=1<<iota
	j=3<<iota
	k      //k=3<<2
	l		//l=3<<3
)

func main() {
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)

	var a string = "abctydrthdghd"
	var b =unsafe.Sizeof(a)
	fmt.Println("a=",a,b)
}
