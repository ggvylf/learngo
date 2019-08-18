package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 100)

	for i := 0; i < 100; i++ {

		select {
		//当case满足多个时候，会随机选一个case执行
		case x := <-ch1:
			fmt.Println("output", x)

		case ch1 <- i:
			fmt.Println("input", i)

		default:
			fmt.Println("do nothing")
		}
	}
}
