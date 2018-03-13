package main

import "fmt"


/*
...展开切片

func MinimumInt(first int, others ...int) （参数1，0个或多个int参数）
MinimumInt(10, 15, 32, 46, 2, 3)  默认10为第一参数，剩余为第二参数
MinimumInt(sliceInt[0], sliceInt[1], sliceInt[2], sliceInt[3], sliceInt[4], sliceInt[5])  //完全同上
MinimumInt(sliceInt[0], sliceInt[0:]...)

 */
func MinimumInt(first int, others ...int) int {
	min := first
	for _, value := range others {
		if value < min {
			min = value
		}
	}

	fmt.Println("min",min)
	return min

}

func Minimum(first interface{}, rest ...interface{}) interface{} {
	min := first
	for _, value := range rest {
		switch value := value.(type) {
		case int:
			if value < min.(int) {
				min = value
			}
		case float64:
			if value < min.(float64) {
				min = value
			}
		case string:
			if value < min.(string) {
				min = value
			}
		}
	}
	fmt.Println("min",min)
	return min
}

func main()  {
	MinimumInt(10, 15, 32, 46, 2, 3)  //1
	var sliceInt = []int{1, 15, 32, 46, 2, 3}
	MinimumInt(sliceInt[0], sliceInt[1], sliceInt[2], sliceInt[3], sliceInt[4], sliceInt[5])  //2
	MinimumInt(sliceInt[0], sliceInt[1:]...)  //3
	fmt.Println("------------------")


	temp :=sliceInt[1:]
	fmt.Println(temp)

	Minimum(10, 15, 32, 46, 2, 3)  //1
	var sliceInt1 = []int{10, 15, 32, 46, 2, 3}
	Minimum(sliceInt[0], sliceInt1[1], sliceInt[2], sliceInt1[3], sliceInt1[4], sliceInt1[5])  //2
	Minimum(sliceInt1[0], sliceInt1[1:])  //3
}
