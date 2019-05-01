package main

import (
	"fmt"
	"time"
)

func recusive() {
	fmt.Println("hello world")
	time.Sleep(time.Second)
	recusive()
}

func main() {
	recusive()
}
