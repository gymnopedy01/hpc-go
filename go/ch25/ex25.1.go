package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()

	//s := []int{}
	//s = append(s, 9)
}

// ch chan int 변수를 받아도 값복사가 이뤄지진 않는다. 레퍼런스만 전달
func square(wg *sync.WaitGroup, ch chan int) {
	defer fmt.Println("DONE")
	n := <-ch
	fmt.Println("Square:", n*n)
	time.Sleep(5 * time.Second)

	wg.Done()
}
