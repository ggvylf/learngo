package main

import (
	"fmt"
	"time"
)

func PutNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func SuSHu(intChan chan int, primeChan chan int, exitChan chan bool) {

	for {
		num, ok := <-intChan
		flag := false

		if !ok {
			break
		}

		for i := 2; i < num; i++ {
			if num%i == 0 {

				flag = false
				break
			}
			flag = true

		}
		if flag {
			primeChan <- num
		}

	}
	fmt.Println("有goroutine取不到数据退出")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	start := time.Now().UnixNano()
	go PutNum(intChan)

	for i := 0; i < 4; i++ {
		go SuSHu(intChan, primeChan, exitChan)
	}

	go func() {

		for i := 0; i < 4; i++ {
			<-exitChan

		}
		end := time.Now().UnixNano()
		fmt.Printf("耗时=%vms\n", (end-start)/1000000)
		close(primeChan)

	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数=", res)
	}
	fmt.Println("主线程退出")

}
