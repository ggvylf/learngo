package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(10 * time.Millisecond)
		select {
		//触发ctx超时
		case <-ctx.Done():
			break LOOP
		default:
		}

	}
	fmt.Println("worker done !")
	wg.Done()

}

func main() {
	//指定超时时间为50ms
	d := 50 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), d)

	wg.Add(1)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("main end")
}
