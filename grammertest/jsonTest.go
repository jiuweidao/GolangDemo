package main

import (
	"encoding/json"
	"fmt"
)

type JsonTest struct {
	X string `json:"x"`
	Y string `json:"y"`
}

func main()  {

	o:=[]JsonTest{}
	a := JsonTest{"a","b"}
	o=append(o,a)
	json_data, _ := json.Marshal(o)

	fmt.Println(string(json_data))
}
