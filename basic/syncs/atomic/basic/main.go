package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomiccandmutex() {
	var a int32 = 3
	//等同于使用mutex
	var locker sync.Mutex
	locker.Lock()
	fmt.Println("use sync.mutex{}")
	a = 5
	fmt.Printf("a=%d,type=%T\n", a, a)
	locker.Unlock()
}

func ops() {
	fmt.Println("use atomic")
	var a int32 = 3
	//从a的地址中获得值
	rst := atomic.LoadInt32(&a)
	fmt.Printf("after load，rst=%d,type=%T\n", rst, rst)
	//把某个值存储到rst的内存地址中，注意没有返回值
	atomic.StoreInt32(&rst, int32(4))
	fmt.Printf("after store，rst=%d,type=%T\n", rst, rst)
	//附加操作等同于相加
	atomic.AddInt32(&rst, 1)
	fmt.Printf("after add,rst=%d,type=%T\n", rst, rst)

	//交换值，并把返回旧值
	old := atomic.SwapInt32(&rst, 1)
	fmt.Printf("after swap,rst=%d,type=%T,old=%d,type=%T\n", rst, rst, old, old)

	//compareandswap，在rst没有被并发修改的情况下，rst的旧值=1的情况下，把rst赋值为3，并返回是否成功的结果
	isswaped := atomic.CompareAndSwapInt32(&rst, 1, 3)
	fmt.Printf("after compareandswap,rst=%d,type=%T,old=%v,type=%T\n", rst, rst, isswaped, isswaped)

}

func main() {
	atomiccandmutex()
	ops()

}
