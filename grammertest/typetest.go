package main

import "fmt"
/*
type_name(expression)
 */
func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32 (sum)/float32(count)
	mean1 := sum/count//float32 (sum)/count异常
	fmt.Printf("mean 的值为: %f\n",mean)
	fmt.Println(mean1)
}