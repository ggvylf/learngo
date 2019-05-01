package main

import "fmt"

func main() {
	//声明变量类型，并赋值
	var i int
	fmt.Println("i=", i)

	//类型推导，根据变量类型自动判断
	var num = 10.11
	fmt.Println("num=", num)
	//省略var 直接声明并初始化
	str := "hello world"
	fmt.Println(str)
}
