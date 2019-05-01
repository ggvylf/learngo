package main

import "fmt"

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

}
