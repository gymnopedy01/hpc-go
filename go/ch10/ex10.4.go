package main

import "fmt"

func main() {
	day := "thursday"

	switch day {
	case "monday", "uesday":
		fmt.Println("월, 화요일은 수업 가는 날 입니다.")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날 입니다.")
	}

}
