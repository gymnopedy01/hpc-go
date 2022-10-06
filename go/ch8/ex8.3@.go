package main

import "fmt"

// 상수값에 코드를 부여합니다.
const (
	Pig int = iota + 6
	Cow
	Chicken
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀", animal)
	} else if animal == Cow {
		fmt.Println("음메", animal)
	} else if animal == Chicken {
		fmt.Println("꼬끼오", animal)
	} else {
		fmt.Println("....", animal)
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)
}
