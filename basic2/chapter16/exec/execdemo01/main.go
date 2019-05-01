package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i < 50; i++ {
		intChan <- i
		fmt.Println("writing", i)
		time.Sleep(time.Microsecond)
	}
	close(intChan)

}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("reading", v)
	}
	exitChan <- true
	close(exitChan)

}

func main() {
	intChan := make(chan int, 50)
	//存储读取channel是否完成，放在main()中判断，避免main()退出
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		ok := <-exitChan
		if !ok {
			break
		}
		fmt.Println("read ok!")
	}
}
