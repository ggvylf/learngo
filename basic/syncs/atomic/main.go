package main

import (
	"fmt"
	"sync/atomic"
)

func intopt() {
	a := int32(1)
	b := int32(-2)
	fmt.Println(atomic.AddInt32(&a, b))
}

func uintopt() {
	a := uint32(1)
	//uint不支持负数，需要通过类型转换
	//方法1 类型转换
	b := int32(-2)
	ub := uint32(b)
	fmt.Println(atomic.AddUint32(&a, ub))

	//方法2 位异或
	ub2 := ^uint32((-(-2) - 1))
	fmt.Println(atomic.AddUint32(&a, ub2))
}

func main() {
	intopt()
	uintopt()
}
