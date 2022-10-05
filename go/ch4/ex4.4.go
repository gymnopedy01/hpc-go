package main

import "fmt"

func main() {

	a := 3
	var b float64 = 3.5

	var c int = int(b) //float64에서 int로 변환
	d := float64(a * c)

	fmt.Println(a, b, c, d)
	fmt.Println(`a ${a}`, a)

	var e int64 = 7
	f := int64(d) * e //float64에서 int64로 변환
	fmt.Println(e, f)

	var g int = int(b * 3) //float64에서 int로 변환
	var h int = int(b) * 3 //float64에서 int로 변환. g와 값이 다릅니다.
	fmt.Println(g, h)

}
