package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once

	//声明一个无参函数，用来被sync.Once执行 该函数只执行一次
	oncefunc := func() {
		fmt.Println("the func run ")
	}

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func(i int) {
			fmt.Println("in goroutine", i)

			//用Do方法执行对应的函数
			once.Do(oncefunc)
			wg.Done()
		}(i)

	}

	wg.Wait()

	fmt.Println("main func end")
}
