package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]float64, 10)
	//全局互斥锁，定义在全局变量
	lock sync.Mutex
)

func JieCheng(n int) {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}

	//加锁,每次写入前家读写锁，解决同时读写竞争的问题
	lock.Lock()
	myMap[n] = float64(sum)
	//解锁
	lock.Unlock()
}

func main() {

	for i := 1; i <= 10; i++ {
		go JieCheng(i)
	}

	time.Sleep(time.Second)

	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("myMap[%v]=%v\n", k, v)
	}
	lock.Unlock()

}
