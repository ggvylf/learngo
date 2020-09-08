package main

import (
	"fmt"
	"strconv"
	"sync"
)

//方法2 使用snyc.Map
// var m = make(map[string]int)
var m = sync.Map{}

//方法1 枷锁
// var lock = sync.Mutex{}

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, value int) {
// 	m[key] = value
// }

func main() {

	wg := sync.WaitGroup{}

	// //超过20就会报错
	// //fatal error: concurrent map writes
	// for i := 0; i < 21; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		lock.Lock()
	// 		key := strconv.Itoa(n)
	// 		set(key, n)
	// 		lock.Unlock()

	// 		fmt.Printf("k:=%v,v=%v\n", key, m[key])
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()

	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k:=%v,v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
