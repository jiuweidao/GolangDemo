package main
import (
	"fmt"
)
func mygoto() {
	i := 0
There:{
	fmt.Println("到达there")
}
	fmt.Println(i)
	i++
	goto There    //跳转到Here
}
func main() {
	mygoto()
}