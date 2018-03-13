package main

import (
	"fmt"
	"time"
)



func main()  {
	go func() {
		fmt.Println("---------")
	}()
	time.Sleep(9000 * time.Second)
}
