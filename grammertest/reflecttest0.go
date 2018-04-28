package main

import (
	"fmt"
)

/*
	nil只能赋值给指针、channel、func、interface、map或slice类型的变量
	interface 类型转换，a,_= b.(int)
    a获取到指定类型的value
 */

func main() {
	animal := 232
	getType(animal)

}

func getType(a interface{}) {
	var b int
	switch t := a.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		b = a.(int)
		fmt.Printf("integer %d %v", t, b) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	case string:

		fmt.Printf("pointer to string %s\n", t) // t has type *int
	}
}
