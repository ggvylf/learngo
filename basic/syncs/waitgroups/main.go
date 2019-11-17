package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	//有1个goroutine 就添加1个wg
	wg.Add(1)
	go func() {
		fmt.Println("timer start=", time.Now())

		end := <-time.After(5 * time.Second)
		fmt.Println("timer end=", end)
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		fmt.Println("ticker start=", time.Now())
		ticker := time.NewTicker(1 * time.Second)
		for i := 0; i < 10; i++ {
			end := ticker.C
			fmt.Println("ticker=", end)
		}
		ticker.Stop()
		fmt.Println("ticker stop=", time.Now())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main func end")
}
