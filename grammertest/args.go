package main



func main() {

	//fmt.Println(powerf3(,480))
}
func powerf3(x float64, n int) float64 {
	ans := 1.0

	for n != 0 {
		ans *= x
		n--
	}
	return ans
}