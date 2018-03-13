package main

import "fmt"


/*

for循环

for {//无限循环
block
}

for booleanExpression {// while循环
block
}

for optionalPreStatement; booleanExpression; optionalPostStatement {// 普通for循环
block
}

for index, value := range string or array or slice {// 普通for range循环，字符串、数组、切片
block
}

for index := range string or array or slice {// 省略value的for range循环，字符串、数组、切片
block
}

for key, value := range map {// 普通for range循环，映射
block
}

for item := range channel{// 通道迭代
block
}
 */
func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {//此时a为临时变量不影响初值
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
}