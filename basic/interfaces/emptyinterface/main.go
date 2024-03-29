package main

import "fmt"

// 空接口，任意类型都实现了空接口 实际在用的时候需要断言
type emptyinf interface{}

// 用法1： 空接口作为函数的参数
func printtype(e emptyinf) {
	fmt.Printf("type is %T\n", e)
}

func main() {
	var e interface{}
	e = 1
	printtype(e)
	e = true
	printtype(e)
	e = [3]int{1, 2, 3}
	printtype(e)
	e = map[int]int{
		1: 1,
		2: 2,
	}
	printtype(e)

	//用法2： 空接口作为map的value
	m := map[int]interface{}{
		1: "hehe",
		2: 222,
		3: []int{1, 2, 3},
	}

	fmt.Println(m)

}
