package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex
var rwlock sync.RWMutex

//互斥所
func testMutex() {
	a := make(map[int]int, 2)
	a[1] = 10
	a[2] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[2] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}

	lock.Lock()
	fmt.Println(a)

	lock.Unlock()

	time.Sleep(time.Second)

}

//读写锁

func testRwmutex() {
	var conut int32
	a := make(map[int]int, 2)
	a[1] = 10
	a[2] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			b[2] = rand.Intn(100)
			time.Sleep(10 * time.Microsecond)
			rwlock.Unlock()
		}(a)
	}

	// for i := 0; i < 100; i++ {
	// 	go func(b map[int]int) {
	// 		rwlock.RLock()
	// 		fmt.Println(b)
	// 		rwlock.RUnlock()
	// 		//原子操作
	// 		atomic.AddInt32(&conut, 1)
	// 	}(a)
	// }

	for {
		go func(b map[int]int) {
			rwlock.RLock()
			//fmt.Println(b)
			time.Sleep(time.Microsecond)
			rwlock.RUnlock()
			//原子操作
			atomic.AddInt32(&conut, 1)
		}(a)
	}

	time.Sleep(3 * time.Second)
	//读取结果
	fmt.Println(atomic.LoadInt32(&conut))

}

func main() {
	// testMutex()
	testRwmutex()
}
