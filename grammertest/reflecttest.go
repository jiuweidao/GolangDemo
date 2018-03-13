package main

import (
	"reflect"
	"fmt"
)
/*
nil只能赋值给指针、channel、func、interface、map或slice类型的变量
 */




type Dogs struct {
	name string `名字`
	age int "年龄"
}

type Animals interface{
	eat()

}

func (dog *Dogs) eat() {
	fmt.Println("i eat bone")
}
func main()  {
	var animal1 Animals =&Dogs{"hsch",23}
	var animal2 Animals
	var dog1 Dogs
	var dog2 Dogs= Dogs{"osf",32}


	a := reflect.TypeOf(animal1)
	b := reflect.ValueOf(animal1)
	c := reflect.TypeOf(dog1)
	dt := reflect.TypeOf(dog2)
	dv := reflect.ValueOf(dog2)

	fmt.Println("reflect.TypeOf(animal1)",a)//*main.Dogs
	fmt.Println("reflect.ValueOf(animal1)",b)//&{hsch 23}
	fmt.Println("b.Interface()",b.Interface())//&{hsch 23}
	//reflact.Value对象可以通过调用Interface()方法，再反射回interface{}对象
	fmt.Println("b.Interface()==animal1",b.Interface()==animal1)//&{hsch 23}
	fmt.Println("reflect.TypeOf(dog1)",c)// main.Dogs
	fmt.Println("reflect.TypeOf(dog2)",dt)//main.Dogs
	fmt.Println("reflect.ValueOf(dog2)",dv)//{osf 32}
	fmt.Println("dv.Interface()",dv.Interface())//{osf 32}
	fmt.Println("dv.Interface()==dog2?",dv.Interface()==dog2)//{osf 32}


	fmt.Println("a.string",a.String())//*main.Dogs
	fmt.Println("a.name",a.Name())//无输出
	for i := 0;i<c.NumField();i++ {
		fild := c.Field(i)
		fmt.Printf("\nc.%s type is:%s",fild.Name,fild.Type)
	}
	
	

	fmt.Println("\nc.string",c.String())//main.Dogs
	fmt.Println("c.name",c.Name())  //Dogs

	if i := reflect.TypeOf(animal2);i==nil{
		fmt.Println("未初始化的接口",i)
	}

	fmt.Println("reflect call")
//	b.MethodByName("eat").Call()




	
}
