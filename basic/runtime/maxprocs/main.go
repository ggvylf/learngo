package main

import (
	"fmt"

	"sync"
)

var wg sync.WaitGroup

func a() {
	for i:=0;i<10;i++ {
		fmt.Printf("func a %d\n",i)
	}
	wg.Done()
}


func b() {
	for i:=0;i<10;i++ {
		fmt.Printf("func b %d\n",i)
	}
	wg.Done()
}



func main() {
	//cpu的逻辑核心数。
	//此时多个goroutine是串行执行的，会看到某个func全部执行完再执行其他的func
	// runtime.GOMAXPROCS(2)
	wg.Add(2)
	go a()
	go b()

	wg.Wait()

}