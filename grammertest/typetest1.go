package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	maptest := []byte(`{"name":"xiao","age":"12,13"}`)

	var object1 interface{}

	if err :=json.Unmarshal(maptest,&object1);err!=nil{

	}else {
		jsonobject := object1.(map[string]interface{})

		fmt.Println(jsonobject)
	}


}
