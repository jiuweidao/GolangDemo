package main

import "fmt"

/*
一、
普通方法与闭包（相似与匿名内部类）的区分
1.普通方法
func  function_name( [parameter list] ) [return_types] {
   函数体
}

2.
func 方法名 返回值（返回方法的方法名 返回值）
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i
   }
}

3.接口
// 定义了有两个方法的接口 I，结构 S 实现了此接口
type I interface {
    Get() int
    Put( int)
}
4.结构体
type Books struct {
	title string
	author string
	subject string
	book_id int
}

二、函数也可以作为形参进行传递
 */

func getNum(a int )(int ,string) {
	return  6,"d"
}
func getsum(a int,getNum func(int)int)  {
	fmt.Println("s=",getNum(a))
}
 
func main()  {
	var x=3
	var y=9
	fmt.Println("x=",x," y=",y)
	x,y=y,x
	fmt.Println("x=",x," y=",y)
	
	num1 :=getNum

	getsum(x,num1)


}
