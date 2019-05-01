package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	//int
	a := int32(1)
	b := int32(-2)
	fmt.Println(atomic.AddInt32(&a, b))

	//uint使用负数
	//方法1 临时变量
	a1 := uint32(1)
	tmp := int32(-2)
	b1 := uint32(tmp)
	fmt.Println(atomic.AddUint32(&a1, b1))

	//方法2 位异或
	a2 := uint32(1)
	b2 := ^uint32((2 - 1)) 
	fmt.Println(b2)
	fmt.Println(atomic.AddUint32(&a2, b2))

}
