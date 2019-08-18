package main

import "fmt"

//单向通道，只能写入
func senttoch(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

//单向通道 ch1只能读取，ch2只能写入
func getfromch(ch1 <-chan int, ch2 chan<- int) {
	//遍历chan中的数据，使用for循环
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go senttoch(ch1)
	go getfromch(ch1, ch2)

	//遍历chan中的数据，使用for-range
	for v := range ch2 {
		fmt.Println(v)
	}
}
