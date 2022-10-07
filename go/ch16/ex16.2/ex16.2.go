package main

import (
	"ch16/ex16.2/publicpkg"
	"fmt"
)

func main() {
	fmt.Println("PI:", publicpkg.PI) //공개되는 상수 접근
	publicpkg.PublicFunc()           //공개되는 함수 호출

	var myint publicpkg.MyInt = 10
	fmt.Println("myint:", myint)

	var mystruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("mystruct:", mystruct)

}
