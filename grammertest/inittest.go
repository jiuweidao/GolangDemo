package main

import "fmt"

var initnum int

func init()  {
	setnum()
}
func main()  {

	fmt.Println(initnum)

}

func setnum()  {

	initnum = 3
}
