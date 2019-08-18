package main

import "fmt"

func main() {

	//声明一个类型为int的chan
	var ch1 chan int
	//初始化带缓冲区的chan，通道长度为1
	ch1 = make(chan int, 1)

	//数据放入chan
	ch1 <- 10

	//从chan中取出数据
	v := <-ch1
	fmt.Println(v)

	//关闭chan
	close(ch1)

}
