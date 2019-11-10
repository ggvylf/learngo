package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask(i int) {
	if i == 5 {
		//出让资源，调整当前goroutine优先级
		runtime.Gosched()
	}

	if i == 9 {
		//查看当前的逻辑cpu个数
		sum := runtime.NumCPU()
		//指定goroutine可用的最大核心数，并返回上一次的配置
		runtime.GOMAXPROCS(sum)
	}

	if i == 3 {
		defer fmt.Printf("the %d goroutine is exit\n", i)
		runtime.Goexit()
		fmt.Println("after exit")
	}
	for j := 0; j < 10; j++ {
		fmt.Printf("im %d task1:%d\n", i, j)
	}

}

func main() {
	fmt.Println("main fun start")
	for i := 0; i < 10; i++ {
		go newTask(i)
		// go newTask2(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("main func end")

}
