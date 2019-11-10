package main

import (
	"fmt"
)

func nobuf() {

	//声明一个类型为int的chan
	var ch1 chan int
	//初始化无缓冲的chan
	ch1 = make(chan int)

	//数据写入chan
	ch1 <- 22
	//从chan中读取
	fmt.Println(<-ch1)
	//关闭chan
	close(ch1)

}

func readfromclosedchan() {
	bools := make(chan bool, 3)
	bools <- true
	bools <- true
	bools <- true
	close(bools)
	for v := range bools {
		fmt.Println(v)
		if string(v) == "nil" {
			fmt.Println("chan is empty")
			return
		}
	}
	bools <- true

}

func withbuf() {
	//声明并初始化，有缓冲，长度为10
	ch2 := make(chan int, 10)

	ch2 <- 1
	ch2 <- 2

	for i := 0; i <= len(ch2); i++ {
		fmt.Println(<-ch2)
	}

}

func main() {

	// nobuf()
	// // withbuf()
	readfromclosedchan()

}
