package main

import (
	"reflect"
	"fmt"
)

/*
1 em必须为initerface类型才可以进行类型断言
	s := "BrainWu"
	if v, ok := s.(string); ok {
		fmt.Println(v)
	}
invalid type assertion: s.(string) (non-interface type string on left)

如果是指针，判断类型，必须通过映射 然后取elem.type 才能判断出是否是和结构体相同类型
以下未正确示范
masssgePoint := &Massage{name:"xiao",age:10}
	valueMassage := reflect.ValueOf(masssgePoint)
	elem := valueMassage.Elem()
	if elem.Type() == reflect.TypeOf(Massage{}){
		fmt.Println("elem.Type() == reflect.TypeOf(Massage{} is true")
		fmt.Println(elem.Type(),reflect.TypeOf(Massage{}))
	}

 */



type Massage struct {
	name string
	age int
}
func main()  {


	/*
	指针类型进行类型判断
	 */
	masssgePoint := &Massage{name:"xiao",age:10}
	valueMassage := reflect.ValueOf(masssgePoint)
	elem := valueMassage.Elem()
	if elem.Type() == reflect.TypeOf(Massage{}){
		fmt.Println("elem.Type() == reflect.TypeOf(Massage{} is true")
		fmt.Println(elem.Type(),reflect.TypeOf(Massage{}))
	}
	fmt.Println()

	if valueMassage.Type() == reflect.TypeOf(Massage{}){
		fmt.Println("valueMassage.Type() == reflect.TypeOf(Massage{} is true")
	}else {
		fmt.Println("valueMassage.Type() == reflect.TypeOf(Massage{} is false")
	}
	fmt.Println(valueMassage.Type(),reflect.TypeOf(Massage{}))
	fmt.Println()

	if masssgePoint,ok := interface{}(masssgePoint).(Massage);ok{
		fmt.Println("interface{}(masssgePoint).(Massage) is true",masssgePoint)
	}else {
		fmt.Println("interface{}(masssgePoint).(Massage) is false",masssgePoint)
	}

	/*
	值类型，进行类型判断
	 */
	masssgeValue := Massage{name:"xiao",age:10}
	valueMassageValue := reflect.ValueOf(masssgeValue)
	if valueMassageValue.Type() == reflect.TypeOf(Massage{}){
		fmt.Println("valueMassageValue.Type() == reflect.TypeOf(Massage{} is true")
		fmt.Println(valueMassageValue.Type(),reflect.TypeOf(Massage{}))
	}else {
		fmt.Println("valueMassageValue.Type() == reflect.TypeOf(Massage{} is false")
		fmt.Println(valueMassageValue.Type(),reflect.TypeOf(Massage{}))
	}

	if masssgevalue,ok := interface{}(masssgeValue).(Massage);ok{
		fmt.Println("interface{}(masssgeValue).(Massage) is true",masssgevalue)
	}
}

