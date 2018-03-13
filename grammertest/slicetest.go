package main
import "fmt"
/*
切片的长度是可变的
未初始化在切片为nil,长度为0，var numbers []int
且数组 可以直接string输出

注：注意区分slice和array
	1.以下都是切片
	slice1:= make([]int,4)
	slice2:= []int{0,12,2,23,3,34}
	2.以下不是切片，不能使用copy和append 即参数里面不能含有非切片
	var temp  []int

一个为nil的slice，除了不能索引外，其他的操作都是可以的，
当你需要填充值的时候可以使用append函数，slice会自动进行扩充。
那么为nil的slice的底层结构是怎样的呢？
 */


func main() {
	slice1:= make([]int,4)
	slice2:= []int{0,12,2,23,3,34}
	var temp  []int
	//for i:=0;i<3;i++{
	//	fmt.Println(slice1[i])
	//}

	fmt.Println(slice2)
	fmt.Println(slice2[1:3])

	temp = append(slice1,6)
	//temp = append(temp,0)

	temp1 := make([]int, len(temp), (cap(temp)))
	copy(temp1,temp)

	fmt.Println(temp)
	fmt.Println(temp1)


	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1,numbers)
	printSlice(numbers1)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}