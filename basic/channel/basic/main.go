package main

import (
	"fmt"
	"time"
)

func nobuf() {

	//声明一个类型为int的chan
	var ch1 chan int
	//初始化无缓冲的chan
	ch1 = make(chan int)

	//从chan中读取
	go func() {
		v, ok := <-ch1
		fmt.Println(v, ok)
	}()
	//数据写入chan
	ch1 <- 22
	//关闭chan
	close(ch1)

}

func readfromclosedchan() {
	bools := make(chan bool, 3)

	bools <- true
	bools <- true
	bools <- false
	close(bools)

	//从关闭的chan中读取数据，把chan中的数据全部读取完成后，返回该类型的零值
	go func() {
		for i := 0; i < 4; i++ {

			v, ok := <-bools
			fmt.Println(v, ok)
		}
	}()
	// 向已经关闭的chan中写入数据会panic
	// bools <- true

}

func closenilchan() {
	// var mychan chan int
	//关闭一个没有初始化的chan会panic
	// close(mychan)
}

func doubleclose() {
	mychan := make(chan int)
	close(mychan)
	close(mychan)
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

func closesignal() {
	mychan := make(chan int, 5)

	go func() {
		for v := range mychan {
			fmt.Println(v)
		}
		fmt.Println("goroutine is end")

	}()

	// for {
	// 	time.Sleep(1 * time.Second)
	// }

	mychan <- 111
	mychan <- 222
	mychan <- 333
	//关闭chan后，会向所有读取的goroutine发送信号不要再阻塞了
	close(mychan)

	defer fmt.Println("out of goroutine")
	time.Sleep(5 * time.Second)

}

func main() {

	nobuf()
	// withbuf()
	// readfromclosedchan()
	// closenilchan()
	// doubleclose()

	// closesignal()
	time.Sleep(2 * time.Second)

}
