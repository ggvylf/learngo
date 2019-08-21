package main

import (
	"fmt"
	"sync"
)

var x map[int]int
var wg sync.WaitGroup
var lock sync.RWMutex

func add(i int) {
	lock.Lock()
	x[1] = i
	fmt.Println("x[1]=", x[1])
	lock.Unlock()
	wg.Done()
}

func read() {
	lock.RLock()
	fmt.Println(x)
	lock.RUnlock()
	wg.Done()
}

func main() {
	x = make(map[int]int, 1)
	x[1] = 1

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go add(i)
		go read()
	}

	wg.Wait()

}
