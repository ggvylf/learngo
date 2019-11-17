package main

import (
	"fmt"
	"sync"
)

var x int = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add(i int) {
	lock.Lock()
	fmt.Printf("the %d goruotine is locked\n", i)

	x += 1

	wg.Done()
	fmt.Printf("the %d gorutine is working, the x=%d\n", i, x)

	fmt.Printf("the %d goruotine is unlocked\n", i)
	lock.Unlock()

}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(i)
	}

	wg.Wait()

	fmt.Println("in func main x=", x)
}
