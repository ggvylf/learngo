package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)

	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		//channel及时没有close，也不会deadline，本身select会继续进行其他的case操作。
		case v := <-intChan:
			fmt.Println("read from intChan=", v)
		case v := <-stringChan:
			fmt.Println("read form stringChan=", v)
		default:
			fmt.Println("全部取到，可以退出了")
			return
		}
	}

}
