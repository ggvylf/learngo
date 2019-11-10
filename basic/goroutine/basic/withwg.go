package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func newTask1(i int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("im %d task1\n", i)

	wg.Done()
}

func newTask2(i int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("im %d task2\n", i)
	wg.Done()
}

func main() {
	fmt.Println("main fun start")
	for i := 0; i < 10000; i++ {
		wg.Add(2)
		go newTask1(i)
		go newTask2(i)
	}
	wg.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println("main func end")

}
