package main


/*

接口
type a interface{
	func1()
	func2()
}
使用结构体实现接口
type b struct {
}
func (b *b) func1{
}

1.如果一个类要继承一个接口，必须实现所有方法
1.注意方法体的一致
3.接口可以内嵌接口


为什么要设计接口？
Go不是传统意义上的面向对象编程语言，
它没有类及其继承的概念。接口是一种契约，实现类型必须满足它，
它描述了类型的行为，规定类型可以做什么。
接口彻底将类型能做什么，以及如何做分离开来，
使得相同接口的变量在不同的时刻表现出不同的行为，这就是多态的本质。
 */
import (
	"fmt"
)
//定义一个接口，接口内有eat和run两个方法
type Animal interface {
	Eat(name string) string
	Run()
}
//定义一个dog实体类，实现Animal接口
type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (dog Dog) Eat(name string) string{
	fmt.Println("I eat bone")
	dog.name = name
	return  dog.name
}

func (dog *Dog) Run() {
	fmt.Println("I run very fast")
}

func (cat *Cat) Eat(name string) string{
	cat.name = name
	fmt.Println(cat.name+"  eat a")
	return  cat.name
}

func (cat *Cat) Run() {
	fmt.Println("I run very be")
}
func main() {//调用接口
	var a = &Dog{} //
	a.Eat("xiaodog")
	a.Run()


	var b Animal
	b = &Cat{}
	catname:=b.Eat("xiaocat")

	switch b.(type) {
		case *Cat:fmt.Println("I am cat ,name:"+ catname)
		case *Dog:fmt.Println("I am dog")
		default:fmt.Println("I don't know who am i")

	}

}