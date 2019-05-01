package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	//使用recover()来捕获panic
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("test() error,err=", err)
		}
	}()

	var myMap map[int]string

	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok!")
		time.Sleep(time.Second)
	}
}
