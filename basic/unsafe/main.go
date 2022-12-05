package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// 把 int类型转换长float64
func intToFloat() {

	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	//输出的结果并不是10
	fmt.Printf("%T %v \n", f, f)
}

// 合理的类型转换
type MyInt int

func convertToMyInt() {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	fmt.Printf("%T %v \n", b, b)
}

// 指针的原子操作
// 通常用在并发的读写操作，为了安全性
// 写的时候另外找地方写，读取的时候把直针指向写的地址上

func ptrAtomic() {

	// 创建一个Pointer类型
	var shareBuffer unsafe.Pointer

	//写函数
	writeDataFn := func() {

		// 初始化数组
		data := []int{}
		for i := 0; i < 5; i++ {
			data = append(data, i)
		}

		// 记录数组对应的地址
		atomic.StorePointer(&shareBuffer, unsafe.Pointer(&data))
	}

	// 读函数
	readDataFn := func() {

		// 从记录中读取数组所在的地址
		data := atomic.LoadPointer(&shareBuffer)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup

	//写入操作
	writeDataFn()
	fmt.Println(shareBuffer)

	// 多个goroutine同时读写
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(100 * time.Microsecond)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(100 * time.Microsecond)
			}
			wg.Done()
		}()

	}

	wg.Wait()
}

func main() {
	intToFloat()
	convertToMyInt()
	ptrAtomic()
}
