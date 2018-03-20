package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (student *Student) String() string {
	return fmt.Sprintf("student:[age:%d ,name:%s]", student.age, student.name)
}

func main() {

	student := &Student{"a", 1}

	fmt.Println(student.String())

}
