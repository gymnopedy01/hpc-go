package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) //context.Background() = 기본컨텍스트 , ctx= 기본컨텍스트에서 캔실이 추가된기능, cancel= cancel 함수
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()

}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): //작업이 끝나면 시그널이 온다.
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")

		}
	}
}
