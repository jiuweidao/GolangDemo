package main

import (
	"fmt"
	"reflect"
)

/*
	Value 的 Call() 方法的参数是一个 Value 的 slice，
	对应的反射函数类型的参数，返回值也是一个 Value 的 slice，
	同样对应反射函数类型的返回值

	反射就是一种检查接口变量的类型和值的机制


 */



func hello() {
	fmt.Println("Hello world!")
}

func helloNum(num int ) {
	fmt.Println("Hello world2!",num)
}
func helloString(num string ) {
	fmt.Println("Hello world2!",num)
}
/*
func main() {
	hl := hello
	hl()
}*/
func main() {
	hl := hello
	fv := reflect.ValueOf(hl)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call([]reflect.Value{})
	fv.Call(nil)

	h2 :=helloNum
	fv2 := reflect.ValueOf(h2)
	params := make([]reflect.Value, 1)                 // 参数
	params[0] = reflect.ValueOf(20)
	fv2.Call(params)

	h3 :=helloString
	fv3 := reflect.ValueOf(h3)
	params3 := make([]reflect.Value, 1)                 // 参数
	params3[0] = reflect.ValueOf("gdfii")
	fv3.Call(params3)






}