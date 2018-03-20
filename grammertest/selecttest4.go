package main
import (
	"fmt"
)

func main() {
	chanCap := 5
	ch := make(chan int, chanCap) //创建channel，容量为5
	for i := 0; i < chanCap; i++ { //通过for循环，向channel里填满数据
		select { //通过select随机的向channel里追加数据
		case ch <- 1:
			fmt.Println("get 1")
		case ch <- 2:
			fmt.Println("get 2")
		case ch <- 3:
			fmt.Println("get 3")
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%v\n", <-ch)
	}
}