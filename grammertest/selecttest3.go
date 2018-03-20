package main

import "fmt"

func mian(){


	a,b :=1,2
// select case must be receive, send or assign recv
	select {
	case a<b:
		fmt.Printf("a<b")
	case b<a:
		fmt.Printf("b<a")
	}
}
