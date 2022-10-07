package main

import "fmt"

func main() {
	var slice []interface{}
	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
	}

	slice = append(slice, 11, 12, 13, 14, 15)
	fmt.Println(slice)
}
