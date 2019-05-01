package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func WriteData(intChan chan int) {
	for i := 0; i < 2000; i++ {
		intChan <- i
		fmt.Println("writeing", i)
		time.Sleep(time.Microsecond)
	}
	close(intChan)

}

func ReadData(intChan chan int, exitChan chan bool) {
	for {
		lock.Lock()
		n, ok := <-intChan
		lock.Unlock()
		if !ok {
			exitChan <- true
			close(exitChan)
			break
		} else {
			fmt.Println("reading", n)
		}

	}

}

func main() {
	intChan := make(chan int, 2000)
	exitChan := make(chan bool, 1)

	go WriteData(intChan)

	for i := 0; i < 8; i++ {
		go ReadData(intChan, exitChan)

	}

	for {
		ok := <-exitChan
		if !ok {
			break
		}

		fmt.Println("read ok")

	}

}
