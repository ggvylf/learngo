package main

import "fmt"

type Test struct {
	x int
	y int
}

func main() {
	var a interface{}
	t := Test{3, 4}
	//空接口可以赋值任意类型
	a = t
	fmt.Println("a=", a)

	b := Test{}
	//类型断言，人为认定a是Test结构体
	// b = a.(Test)

	//带判断的类型断言
	b,ok:=a.(Test)
	if ok {
		fmt.Println("Ok!")
	}else {
		fmt.Println("not ok")
	}


	fmt.Println("b=", b)

}
