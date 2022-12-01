package main

import (
	"fmt"
	"sync"
)

var x int = 0
var wg sync.WaitGroup
var lock sync.Mutex

// 非线程安全的 最后的结果可能不是5000 goroutine是并行的
func addThreadUnSafe(i int) {

	x += 1

	fmt.Printf("the %d gorutine is working, the x=%d\n", i, x)

}

// 线程安全的
func addThreadSafe(i int) {
	defer func() {
		lock.Unlock()
	}()

	lock.Lock()
	fmt.Printf("the %d goruotine is locked\n", i)

	x += 1

	wg.Done()
	fmt.Printf("the %d gorutine is working, the x=%d\n", i, x)

	fmt.Printf("the %d goruotine is unlocked\n", i)

}

func main() {

	// for i := 0; i < 5000; i++ {
	// 	go addThreadUnSafe(i)
	// }
	// time.Sleep(3 * time.Second)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go addThreadSafe(i)
	}

	wg.Wait()

	fmt.Println("in func main x=", x)
}
