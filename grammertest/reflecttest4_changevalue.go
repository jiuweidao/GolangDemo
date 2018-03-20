package main

import (
	"reflect"
	"fmt"
)

func main()  {
	x := 2
	value :=reflect.ValueOf(&x)
	d := value.Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)

	d = 4
}
