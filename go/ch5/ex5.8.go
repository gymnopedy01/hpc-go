package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil { //에러 발생시
		fmt.Println(err)       //에러출력
		stdin.ReadString('\n') // 표준 입력 스트림 지우기
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil { //에러 발생시
		fmt.Println(err) //에러출력
	} else {
		fmt.Println(n, a, b)
	}
}
