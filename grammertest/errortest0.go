package main
import (
	"errors"
	"fmt"
)

func requireDual(n int) (int, error) {
	if n&1 == 1 {
		return -1, errors.New("您输入不是双数") //生成一个简单的 error 类型
	}

	return n, nil
}

func main() {
	if result, err := requireDual(101); err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("结果：", result)
	}
}