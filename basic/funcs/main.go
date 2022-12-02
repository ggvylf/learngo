package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(a int, b ...int) (sum int) {
	sum = a
	for _, v := range b {
		sum += v
	}

	return
}

func coucat(a, b int) int {
	//匿名函数
	return func(a1, b1 int) int {
		return a1 + b1
	}(a, b)
}

func main() {
	sum := add(1, 2, 3, 4)
	fmt.Println(sum)

	fmt.Println(coucat(3, 4))

	// 多返回值函数
	a, b := returnMultiValues()
	fmt.Println(a, b)

	// 函数是一等公民
	// 计算函数执行时间的函数
	tsSF := timeSpent(slowFunc)
	tsSF(10)

	//函数可变参数
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2))

	//defer函数
	testDefer()

}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)

}

// 计时函数
// 创建具体执行任务的函数
func slowFunc(op int) int {
	time.Sleep(1 * time.Second)
	return op
}

// 计算函数运行的时间
// 把slowFunc 这个函数包了一下
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()

		// 执行传入的函数
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

// 可变参数
func Sum(ops ...int) int {
	sum := 0
	for _, op := range ops {
		sum += op
	}
	return sum
}

// defer函数
// defer函数先执行，后panic在
func useInDefer() {
	fmt.Println("print by defer()")
}

func testDefer() {
	defer useInDefer()
	fmt.Println("this is test defer func")
	panic("this is test defer func panic")
}
