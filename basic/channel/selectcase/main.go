package main

import (
	"fmt"
	"time"
)

func main() {

	chanA := make(chan int, 5)
	chanB := make(chan int, 4)
	chanC := make(chan int, 4)

OUTER:
	for {
		//同时相应多个通道操作
		select {
		case chanA <- 1:
			fmt.Println("　write to chanA")
		case chanB <- 1:
			fmt.Println("　write to chanB")
		case chanC <- 1:
			fmt.Println("　write to chanC")
		default:
			fmt.Println("all tatsk end")
			break OUTER
		}
	}

	time.Sleep(5 * time.Second)

	fmt.Println("main func end")
}
