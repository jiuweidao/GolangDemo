package main

import "fmt"

type Foo struct {
	bar string
}
func main() {
	list := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	list2 := make([]*Foo, len(list))
	for i, value := range list {
		list2[i] = &value
	}
	/*
	此处每次改变的是value地址对应的值，每次取地址给list2填充值的时候，都是同一个地址
	 */
	fmt.Println(list[0], list[1], list[2])
	fmt.Println(list2[0], list2[1], list2[2])
}
