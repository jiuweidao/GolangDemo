package main

import "fmt"

/*
类型断言语法，x为接口变量，T为任意的类型
x.(T)
Comma-ok断言的语法是：
value, ok := element.(T)。
element必须是接口类型的变量，T是普通类型。
如果断言失败，ok为false，value 为[],否则ok为true并且value为变量的值。
简化版：
value := element.(T)，若断言失败，程序异常
注：
同包下的命名是相同
 */
type Html []interface{}

func main(){

	html := make(Html, 5)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = "style"
	html[4] = "head"

	if value,ok := html[1].([]byte);!ok {
		fmt.Println("匹配失败后的异常")
		fmt.Println(value)

	}

	for index, element := range html {
		if value, ok := element.(string); ok {
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.([]byte); ok {
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		}
	}


	for index, element := range html {
		switch value := element.(type) {
		case string:
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		case []byte:
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		case int:
			fmt.Printf("invalid type\n")
		default:
			fmt.Printf("unknown type\n")
		}
	}



}
