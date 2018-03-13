package main

import "fmt"
/*
1.map集合需初始化，方式如下
声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type
使用 make 函数
map_variable := make(map[key_data_type]value_data_type)

2.captial, ok := countryCapitalMap["India"]
captical为该key 对应的value

3.支持直接打印输出
 */
func main() {
	var countryCapitalMap map[string]string
	/* 创建集合 */
	//若不进行下面的初始化则assignment to entry in nil map
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["India"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if(ok){
		fmt.Println("Capital of United States is", captial)
	}else {
		fmt.Println("Capital of United States is not present")
	}

	delete(countryCapitalMap,"France");
	fmt.Println(countryCapitalMap)
}