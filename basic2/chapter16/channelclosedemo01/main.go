package main

import "fmt"

func main() {
	intChan := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}
	close(intChan)

	// 不能使用for循环，因为每次遍历len(intChan)是变化的
	// for i := 0; i < len(intChan); i++ {
	// 	n := <-intChan
	// 	fmt.Println(i, n)
	// }

	//for-range方式，管道必须已经关闭，否则会deadline，但是数据已经是全部取出但是无法退出循环。channel没有index的概念，只有value。
	for v := range intChan {
		fmt.Println("v=", v)
	}
}
