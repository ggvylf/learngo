package main

import (
	"fmt"
	"time"
)

func HelloWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world", i)
		time.Sleep(time.Second)

	}

}

func main() {
	//开启一个goroutine
	go HelloWorld()
	for i := 0; i < 10; i++ {
		fmt.Println("hello main", i)
		time.Sleep(time.Second)
	}

}
