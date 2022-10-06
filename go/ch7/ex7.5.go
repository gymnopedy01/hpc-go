package main

import "fmt"

func Divide(a, b int) (result int, success bool) { //함수선언
	if b == 0 {
		result = 0
		success = false
		return //제수가 0일때 반환
	}
	result = a / b
	success = true
	return //제수가 0이 아닐 때 반환
}

func main() {
	c, success := Divide(9, 3) //제수가 0이 아닌 경우
	fmt.Println(c, success)
	d, success := Divide(9, 0) //제수가 0인 경우
	fmt.Println(d, success)
}
