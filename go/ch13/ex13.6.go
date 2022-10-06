package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32
	Score float64
}

func main() {
	user := User{23, 77.2}

	//Age 4바이트, Score 8바이트
	fmt.Println(unsafe.Sizeof(user.Age) + unsafe.Sizeof(user.Score))

	//하지만 16바이트로 출력, 메모리 정렬 문제 때문
	fmt.Println(unsafe.Sizeof(user))
}
