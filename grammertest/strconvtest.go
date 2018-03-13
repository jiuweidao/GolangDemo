package main

import (
	"strconv"
	"fmt"
)

func main()  {
	x,error :=strconv.ParseFloat("-99.7",64)
	fmt.Printf("%-20T %6v %v\n",x,x,error)

	y,error :=strconv.ParseInt("9",10,0)
	fmt.Printf("%-20T %6v %v\n",y,y,error)

	z,error :=strconv.Atoi("71309")
	fmt.Printf("%-20T %6v %v\n",z,z,error)
}
