package main
import "fmt"


/*
若为空指针 打印输出ptr为，否则为ptr地址  var  ptr *int=7  错误
 */
func main() {

	var a= 7
	var  ptr *int =&a

	fmt.Printf("ptr 的值为 : %x\n", ptr  )
}
