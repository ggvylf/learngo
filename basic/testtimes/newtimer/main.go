package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器
	timer := time.NewTimer(3 * time.Second)

	//从定时器的channel中读取数据
	time := <-timer.C
	fmt.Println(time)
	fmt.Println("prog run ok!")
}
