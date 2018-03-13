package main

import "fmt"

var balance [10] float32
var balance1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

var identifier []int
var slice1 []int = make([]int, 10)
//slice1 := make([]type, len)
//make([]T, length, capacity)

func GetNum(nums []int)  {
	if len(nums)>0{
		fmt.Printf("getnum %d",len(nums))
	}
}
func main()  {
	s :=[] int {1,2,3 }
	s1 := balance[:]
	fmt.Println(s,s1)



	a := []string{"a","b","c"}
	b := []string{"c","d"}
	//var fu =GetNum()
	GetNum([]int{2,3})


	a = append(a,b[1:2]...)
	fmt.Println(a)



	}