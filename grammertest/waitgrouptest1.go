package main


import (
	"fmt"
	"sync"
	_"time"

)

func main() {
	var wg sync.WaitGroup

	for i := 0; i > 5; i = i + 1 {
		wg.Add(1)
		go EchoNumber(i)
	}

	wg.Wait()
	//time.Sleep(time.Millisecond * 12)
}

func EchoNumber(i int) {
	//time.Sleep(3e9)
	fmt.Println("_________________---",i)
}
