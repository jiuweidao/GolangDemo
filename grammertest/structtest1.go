package main

import "fmt"
/*
聚合与嵌入区分以及命名冲突
 */
type Person struct {
	Title     string
	Forenames []string
	Surname   string
}

type Author1 struct {
	Names    Person // 聚合
	Title    []string
	YearBorn int
}

type Author2 struct {
	Person   // 嵌入，匿名结构体字段
	Title    []string
	YearBorn int
}

func Test_embedAndAggregation() {
	author1 := Author1{Person{"Person", []string{"X", "U"}, "ZHE"}, []string{"XU", "ZHE"}, 19911115}
	author1.Names.Title = "Male" // 聚合结构体的字段访问
	fmt.Println(author1)

	author2 := Author2{Person{"Person", []string{"M", "A"}, "ZHUOJUN"}, []string{"MA", "ZHUO", "JUN"}, 19920227}
	author2.Surname = "ZHUOJUN2"// 聚合结构体的字段访问，不存在冲突
	author2.Person.Title = "Female"// 聚合结构体的字段访问，存在冲突
	fmt.Println(author2)
}