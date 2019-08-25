package main

import (
	"fmt"
	"sync"
)

var x int = 10
var wg sync.WaitGroup
var lock sync.Mutex

func add(i int) {
	lock.Lock()
	defer lock.Unlock()
	x = i + 1

	wg.Done()
	fmt.Printf("the %d times in func x=%d\n", i, x)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(i)
	}

	wg.Wait()
	fmt.Println("in func main x=", x)
}
