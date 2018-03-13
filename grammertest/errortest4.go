package main

import (
	"github.com/anker-dev/infra/log"
	"fmt"
)

func getOne(i int)error{

	if i>0 {
		return log.Error("i>0 wrong")
	}else {
		return  nil
	}

}

func main()  {
	err :=getOne(2)
	fmt.Printf("%s",err.Error())
}