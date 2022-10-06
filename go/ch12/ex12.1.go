package main

import "fmt"

func main() {

	//var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 262.2}

	//var t = [5]int{1: 10, 3: 30}

	t := [...]int{10, 20, 30}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

}
