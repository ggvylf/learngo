package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

//单向通道，只能写入
func senttoch(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	//写入数据完成后，关闭chan，此时只能读不能写入
	close(ch1)
	wg.Done()
}

//单向通道 ch1只能读取，ch2只能写入
func getfromch(ch1 <-chan int, ch2 chan<- int) {
	fmt.Println("start get")
	//遍历chan中的数据，使用for循环
	for {
		tmp, ok := <-ch1
		//已经读完ch1中的数据
		if !ok {
			break
		}
		ch2 <- tmp * tmp
		fmt.Println("tmp=", tmp)
	}

	wg.Done()
	once.Do(func() {
		close(ch2)
		fmt.Println("ch2 is closed")
	})

}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	wg.Add(3)
	// 1个生产者
	go senttoch(ch1)

	//2个消费者
	go getfromch(ch1, ch2)
	go getfromch(ch1, ch2)

	//遍历chan中的数据，使用for-range
	for v := range ch2 {
		fmt.Println(v)
	}
	wg.Wait()
}
