package main

import (
	"time"
	"fmt"
	//"strconv"
	"strconv"
	"reflect"
)

func main() {

	t1 := time.Now()
	//print(t1)
	t2 := t1.AddDate(0, -2, 0)
	fmt.Println(t1)
	fmt.Println(t2)

	timestamp := time.Now().Unix()

	fmt.Println(timestamp)

	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm)

	tm1 := time.Unix(timestamp, 0)
	fmt.Println(tm1)
	fmt.Println(tm1.Format("20060102"))

	now, _ := strconv.Atoi(time.Now().AddDate(0, 0, 1).Format("20060102"))

	fmt.Println(t1.Format("20060102"))
	fmt.Println(now)

	//t1.Format("20060102")

	//unixInt, err := strconv.ParseInt("2016", 10, 64);


	a := time.Now()
	b := a.AddDate(0, 0, 0)
fmt.Println(b)
	c:=b.Unix()
	fmt.Println(c)

	the_time := time.Date(2018, 4, 18, 0, 0, 0, 0, time.Local)
	unix_time := the_time.Unix()
	fmt.Println("方法一 时间戳：", unix_time, reflect.TypeOf(unix_time))

	the_time = time.Date(2018, 4, 28, 0, 0, 0, 0, time.Local)
	unix_time = the_time.Unix()
	fmt.Println("方法一 时间戳：", unix_time, reflect.TypeOf(unix_time))

}
