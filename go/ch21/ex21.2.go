package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("반드시 호출됩니다.") //지연 수행될 코드
	defer f.Close()                 //지연수행될코드
	defer fmt.Println("파일을 닫았습니다.") //지연수행될코드

	fmt.Println("파일에 Hello Worl를 씁니다.")
	fmt.Fprintln(f, "Hello World") //파일에 텍스트를 씁니다.

}
