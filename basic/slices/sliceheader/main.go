package main

import (
	"fmt"
	"unsafe"
)

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func main() {
	// len为0 cap为2
	s := make([]int, 0, 2)
	fmt.Println("s=", s)

	// 打印sliceheader
	sh := (*SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("s sliceheader=", sh)

	doSomething(s)
	fmt.Println("after s=", s)

	// 方法1 重新切片
	fmt.Println("after s[:1]=", s[:1])

	// 方法2 将append的结果保存起来
	s2 := doSomething2(s)
	fmt.Println("s2=", s2)

}

// 函数中复制的是slice的sliceheader，创建了sliceheader的新副本
func doSomething(a []int) {
	a = append(a, 1)
	fmt.Println("a=", a)

	// 打印sliceheader
	sh := (*SliceHeader)(unsafe.Pointer(&a))
	fmt.Println("a sliceheader=", sh)
}

func doSomething2(a []int) []int {
	a = append(a, 1)
	fmt.Println("a=", a)

	// 打印sliceheader
	sh := (*SliceHeader)(unsafe.Pointer(&a))
	fmt.Println("a sliceheader=", sh)
	return a
}
