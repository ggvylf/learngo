package main

import (
	"fmt"
	"runtime"
	"sync"
)

//全局声明
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello world")
	//计数器-1
	wg.Done()
}

func hello2() {
	fmt.Println("hello world2")
	//计数器-1
	wg.Done()
}

func main() {
	//制定P的个
	runtime.GOMAXPROCS(1)

	//计数器+1
	wg.Add(1)
	//使用go关键字开启goroutine
	go hello()

	//开启多个goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go hello2()

	}

	//匿名函数
	//注意闭包传参的问题
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {

			fmt.Println("hello =", i)
			wg.Done()
		}(i)
	}

	fmt.Println("hello main func")
	//等待计数器为空后结束
	wg.Wait()

}
