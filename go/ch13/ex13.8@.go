package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8  //1바이트
	C int8  //1바이트
	B int64 //4바이트
	D int8  //2바이트
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
	fmt.Println(unsafe.Sizeof(user.A))
	fmt.Println(unsafe.Sizeof(user.C))
	fmt.Println(unsafe.Sizeof(user.B))
	fmt.Println(unsafe.Sizeof(user.D))

	fmt.Println(&user.A)
	fmt.Println(&user.C)
	fmt.Println(&user.B)
	fmt.Println(&user.D)

}
