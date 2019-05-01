package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var i int = 1
	fmt.Println("i=", i)

	var j int8 = -128
	fmt.Println("j=", j)

	var k uint16 = 255
	fmt.Println("k=", k)

	fmt.Printf("%T,%d\n", k, unsafe.Sizeof(k))

	var n1 float32 = -123.12345678910083
	var n2 float64 = -123.12345678910083
	var n3 = -123.12345678910083
	n4 := 2.0e8
	fmt.Println(n1, n2, n4)
	fmt.Printf("%T\n", n3)
}
