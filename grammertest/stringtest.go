package main

import (
	_ "fmt"
	"fmt"
	"strings"
)


func main()  {
	s :=  "sihch中文"

	for index,char :=range s  {
		fmt.Printf("%d%c\t",index,char)
	}

	fmt.Println("\ns index输出")

	for i:=0;i<len(s) ; i++ {
		fmt.Printf("%d%c\t",i,s[i])
	}


	fmt.Println("\nrun输出")
	chars:=[]rune(s)
	for i:=0;i<len(chars) ; i++ {
		fmt.Printf("%d%c\t",i,chars[i])
	}

	fmt.Println(strings.Contains(s,"中文"))
	fmt.Println(strings.ContainsAny(s,"汉字"))
	fmt.Println("strings.Index(s,'文'",strings.Index(s,"文"))
	fmt.Println("strings.Index(s,''",strings.IndexByte(s,''))
	fmt.Println("strings.Index(s,''",strings.IndexRune(s,rune('s')))
	fmt.Println(strings.IndexRune("widuu", rune('w'))) //0

}