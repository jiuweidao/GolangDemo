package main

import "fmt"
/*
一、
for k,v := range b
遍历切片，key是index 是value
如果只想输出value而不想要key，

若只含一个for i := range b
默认是索引

for k ,_:= range b 保留keyy遗弃value
for _,v:= range b 与上面相反
像下面这么写编译会出错，
因为go语言对于“声明而未被调用”的变量在编译阶段会报错，
但是可以用“_”下划线来丢弃

二、可变参数函数
pairs ...Exchanger


 */
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for k, v := range a {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}

	fmt.Println("v------------------")
	b := [5]int{1, 2, 3, 4, 5}
	for v ,_:= range b {
		fmt.Println("value:", v)
	}
}