package main

import "fmt"

func main() {
	var x1 int8 = 34
	var x2 int16 = 34
	var x3 uint8 = 34
	var x4 uint16 = 34

	fmt.Printf("^%d = %5d,\t %08b\n", x1, ^x1, uint8(^x1))
	fmt.Printf("^%d = %5d,\t %016b\n", x2, ^x2, uint8(^x2))
	fmt.Printf("^%d = %5d,\t %08b\n", x3, ^x3, uint8(^x3))
	fmt.Printf("^%d = %5d,\t %016b\n", x4, ^x4, uint8(^x4))

	var x5 uint16 = 10
	var x6 uint16 = 2
	fmt.Printf("^%d = %5d,\t %16b\n", x5, x6, x5&^x6)

}
