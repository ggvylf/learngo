//闭包
//一个函数内部使用了外部的变量

package main

import "fmt"

func Adder() func(int) int {
	//初始化外部变量
	var n int

	return func(x int) int {
		//修改外部变量的值
		fmt.Printf("n is %d\n", n)
		n += x
		fmt.Printf("new n is %d\n", n)
		return n

	}

}

func main() {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))
}
