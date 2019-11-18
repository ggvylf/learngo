package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	mych := make(chan bool)

	//声明一个无参函数，用来被sync.Once执行
	oncefunc := func() {
		fmt.Println("once run ok")
		mych <- true
	}

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func(i int) {
			fmt.Println("in goroutine", i)
			once.Do(oncefunc)
			wg.Done()
		}(i)

	}

	<-mych
	wg.Wait()

	fmt.Println("main func end")
}
