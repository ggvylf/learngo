package main

import "fmt"

func main() {
	var intChan chan int
	//channel的容量是固定的
	intChan = make(chan int, 3)

	numin1 := 10
	numin2 := 20

	//写入数据到channel
	intChan <- numin1
	intChan <- numin2

	fmt.Printf("intChan的长度=%v,容量=%v\n", len(intChan), cap(intChan))

	//从channel中读取数据
	numout1 := <-intChan

	fmt.Printf("numout1的值是%v,intChan的长度=%v,容量=%v\n", numout1, len(intChan), cap(intChan))
	numout2 := <-intChan
	fmt.Printf("numout2的值是%v,intChan的长度=%v,容量=%v\n", numout2, len(intChan), cap(intChan))

	numout3 := <-intChan
	fmt.Printf("numout2的值是%v,intChan的长度=%v,容量=%v\n", numout3, len(intChan), cap(intChan))
}
