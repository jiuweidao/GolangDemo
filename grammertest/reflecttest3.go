package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
		和 reflect.Type 类似,
	reflect.Value 也满足 fmt.Stringer 接口,
	但是除非 Value 持有的是字符串, 否则 String 只是返回具体的类型.
	 */
	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.Int())    //"3)"
	fmt.Println(v.String()) // NOTE: "<int Value>"

	v1 := reflect.ValueOf("qwdw")
	fmt.Println(v1)          // "3"
	fmt.Printf("%v\n", v1)   // "3"
	//fmt.Println(v1.Int())    // NOTE: "<int Value>"call of reflect.Value.Int on string Value
	fmt.Println(v1.String()) // NOTE: "<int Value>"


	/*
	一个 reflect.Value 和 interface{} 都能保存任意的值.
	所不同的是, 一个空的接口隐藏了值对应的表示方式和所有的公开的方法,
	即
		reflect.ValueOf(3)v返回的是一个可见内部值以及方法的结构体值、
		而value.interface()返回的是一个隐藏了其内部信息的接口
	因此只有我们知道具体的动态类型才能使用类型断言来访问内部的值(就像上面那样),
	对于内部值并没有特别可做的事情.
	相比之下, 一个 Value 则有很多方法来检查其内容, 无论它的具体类型是什么.

	所有通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的
	我们可以通过调用reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的Value。
	 */
	v2 := reflect.ValueOf(3) // a reflect.Value
	x := v2.Interface()      // an interface{}

	i := x.(int)            // an int
	fmt.Printf("%v\n", v2.Int())
	fmt.Printf("%v\n", x)
	fmt.Printf("%d\n", i)
}