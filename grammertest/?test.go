package main

import "fmt"

func main() {
	//"INSERT INTO `device` VALUES ('AET5550815100001', 'Z6111');"
	for i := 0; i < 188; i++ {

		fmt.Printf("INSERT INTO `device` VALUES ('AET5550815%d', 'Z6111');\n", i+100001)
	}
}
