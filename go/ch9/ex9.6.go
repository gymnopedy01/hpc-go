package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are Young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age) //age 값에 접근가능
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	//fmt.Println("Your age is ", age) //Error - age 는 소멸되었습니다.

}
