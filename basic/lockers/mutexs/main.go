package main

import (
	"fmt"
	"sync"
)

var x = 10
var wg sync.WaitGroup
var lock sync.Mutex

func add(i int) {
	lock.Lock()
	x = i + 1
	fmt.Println("x=", x)
	lock.Unlock()
	wg.Done()

}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(i)
	}

	wg.Wait()
	fmt.Println(x)

}
