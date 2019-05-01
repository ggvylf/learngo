package main

import (
	"context"
	"fmt"
	"time"
)

func test() {
	//匿名函数，返回值是单向channel
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("context done")
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	//初始化ctx
	ctx, cancel := context.WithCancel(context.Background())
	//关闭ctx
	defer cancel()

	//调用gen()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
}

func main() {

	test()
	time.Sleep(5 * time.Second)
	fmt.Println("all done")
}
