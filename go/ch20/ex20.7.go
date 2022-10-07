//425페이지

package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (S *Student) String() string {
	return fmt.Sprintf("Student Age:%d", S.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	// var s *Student = &Student{15}

	// PrintAge(s)

	s := &Student{15}

	PrintAge(s)
}
