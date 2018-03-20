package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	/*fmt.Println("Commencing countdown.")
	tick := time.Tick(2 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		j := <-tick
		fmt.Println(j)
	}
	launch()*/



	abort := make(chan int)

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte

		fmt.Println("开始发送数据")
		abort <- 1
		time.Sleep(time.Millisecond * 100)
	}()
	//time.Sleep(time.Millisecond * 1000)
	fmt.Println("Commencing countdown.  Press return to abort.")

	select {
		case <-time.After(10 * time.Second):
		//j := <-tick
		fmt.Println("come here")
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	/*default:
		fmt.Println("default")*/
	}
	launch()
	//<-abort
	time.Sleep(time.Millisecond * 100)
}

func launch()  {
	fmt.Println("火箭发射")
}



