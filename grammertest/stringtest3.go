package main

import (
	"fmt"
)

func main()  {
	a:=[]string{"a","b","c"}
	var b string
	for i,c :=range a{

		if i<len(a)-1{
			b+=c+","
		}else {
			b+=c
		}
	}

	fmt.Println(b)


}
