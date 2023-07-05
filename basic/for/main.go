package main

import "fmt"

func main() {

	test1()
	test2()

}

// for是循环变量，共用一个地址
// go1.21开始使用GOEXPERIMENT=loopvar来改为迭代地址
// go1.22会默认启用

// 指针的陷阱
func test1() {

	s := []int{1, 2, 3}
	d := []*int{}

	for _, v := range s {
		// v指向的内存地址存放的是循环的最后一个元素
		// 使用局部变量可以避免
		// v := v
		fmt.Printf("range中的v=%#v,ptr=%#v\n", v, &v)
		d = append(d, &v)

	}
	fmt.Println("d=", d)
}

// 闭包的陷阱
func test2() {
	var prints []func()
	for _, v := range []int{1, 2, 3} {
		// v指向的内存地址存放的是循环的最后一个元素
		// 使用局部变量可以避免
		// v := v
		fmt.Printf("range中的v=%#v,ptr=%#v\n", v, &v)
		prints = append(prints, func() { fmt.Println(v) })
	}

	for _, print := range prints {
		print()
	}
}
