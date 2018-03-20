package main

import (
	"fmt"
	"reflect"
)


type Stringer interface {
	String() string
}

type Book struct {
	name string
}

func (book Book) String() string  {
	return  book.name
}





func main(){
	var book1 Stringer = Book{"go语言"}
	sv, ok := book1.(Stringer)
	if ok {
		fmt.Printf("v implements String(): %s\n", book1.String()) // note: sv, not v
		fmt.Printf(sv.String()) // note: sv, not v
	}

	var book2 Stringer

	typebook2 := reflect.TypeOf(book2)
	fmt.Println(typebook2)


}
