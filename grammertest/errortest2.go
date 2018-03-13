package main
import (
	"fmt"
)

type dualError struct {
	Num     int
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("参数不正确，因为\"%d\"不是双数", e.Num)
}

func requireDual1(n int) (int, error) {
	if n&1 == 1 {
		return -1, dualError{Num: n}
	}

	return n, nil
}

func main() {
	if result, err := requireDual1(101); err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("结果：", result)
	}
}