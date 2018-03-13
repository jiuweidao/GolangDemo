package main

import "fmt"

//接口就是一些了具有相同结果的结构体指针一样的东西，可以调用形式一样的结构体实体
type Phone interface { //接口类型，定义和结构体差不多
	call()
}

type NokiaPhone struct { //定义结构体
}
type IPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia,I can call you!")
}
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

//定义一个错误结构体
type DivideError struct {
	dividee int
	divider int
}

// 实现 	`error` 接口
func (de *DivideError) Error() string {
	//下面是一个字符串，为了分行写，还有在sprintf里面会简单点
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)

	/*
	/* 用sprintf函数将x和message输出到字符数组buffer中
		sprintf(buffer, "%s%d", message, x);
	 */
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,//给属性赋值用：
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	//错误研究

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}


