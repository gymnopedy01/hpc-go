package main

import "fmt"

func main() {
	str1 := "Hello\t'tWorld'\n"

	str2 := `Go is "aweesome"!\nGo is simple and \t powerful'`
	fmt.Println(str1)
	fmt.Println(str2)
}
