package main

import (
	"fmt"
	"time"
)

func lenandcap() {
	myChan := make(chan int, 3)
	fmt.Printf("myChan len=%d\n", len(myChan))
	fmt.Printf("myChan len=%d\n", cap(myChan))
}

func readandwriter() {
	myChan := make(chan int, 3)
	for i := 0; i < 3; i++ {
		myChan <- i
	}

	// 不能向已满的chan中写入数据，会panic
	// myChan <- 4

	<-myChan
	<-myChan
	<-myChan
	//不能从空chan中读取数据，会panic
	// <-myChan

}

func main() {
	// lenandcap()
	readandwriter()
	time.Sleep(3 * time.Second)
	fmt.Println("man func is end")

}
