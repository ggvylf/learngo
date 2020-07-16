package main

import (
	"fmt"
	"sync"

)


 var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello",i)

}

func main() {
	//调用外部函数
	//通过for循环来控制goroutine的数量
	// for i:=0;i<10;i++ {
	// 	go hello(i)
	// }


	//匿名函数错误写法
	//没有把i传入匿名函数内部。导致从外部获取i的值，这是闭包，goroutine在执行的时候，i的值可能来不及变化。
	// for i:=0;i<10;i++ {
	// 	go func(){
	// 		fmt.Println("hello",i)
	// 	}()

	// }



	//匿名函数
	// for i:=0;i<10;i++ {
	// 	go func(i int){
	// 		fmt.Println("hello",i)
	// 	}(i)

	// }


	//使用sync.WaitGroup实现主进程不退出
	//匿名函数

	for i:=0;i<10;i++ {
		//计数器+1
		wg.Add(1)
		
		go func(i int){
			fmt.Println("hello",i)
			//计数器-1
			wg.Done()
		}(i)
		

	}

	//等待计数器归0
	wg.Wait()

	//保证goroutine执行完之前 main函数不退出
	// time.Sleep(1* time.Second)

}