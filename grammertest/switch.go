package main

import "fmt"
/**
switch不需要break
switch分支，Go语言的case不会自动向下贯穿，需使用fallthrough关键字
 */
func main() {
	var x interface{}

	//var a int =10
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T",i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型" )
	default:
		fmt.Printf("未知型")
	}
}