package main

import "fmt"

func main() {
	fmt.Println("print test1() return v=", test1())
	fmt.Println("print test1() return v=", test2())
	fmt.Println("print test1() return v=", test3())
	fmt.Println("print test1() return v=", test4())
}

// 执行顺序
// return最先执行->return负责将结果写入返回值中->接着defer开始执行一些收尾工作->最后函数携带当前返回值退出
// 函数的返回值取决于return
// 函数的返回值只要提前声明，defer的操作不会影响函数的返回值

// 子函数的defer执行以后 才会执行main()函数里的fmt.Println

// 函数初始化v=0 main函数v=0 defer里 v=0
// 0 0
func test1() (v int) {
	defer fmt.Println("print test1 defer v=", v)
	return v
}

// 匿名函数
// 初始化v=0 main函数v=3 defer里入栈的是函数，不是值 执行的时候v=3
// 3 3
func test2() (v int) {
	defer func() {
		fmt.Println("print test2 defer v=", v)
	}()
	return 3
}

// 函数初始化 v=0  函数内部v=3 main函数v=4 defer处理以后v=4
// 0 4
func test3() (v int) {
	// 提前声明返回值
	// v = 3
	defer fmt.Println("print test3 defer v=", v)
	v = 3
	return 4
}

// 函数初始化v=0  mian函数v=5 defer入栈函数v=0,执行的时候v=5
// 0 5
func test4() (v int) {
	defer func(n int) {
		fmt.Println("print test4 defer v=", n)
	}(v)
	return 5
}
