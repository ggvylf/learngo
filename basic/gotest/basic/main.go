package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sum(a int) int {
	sum := 0
	for i := 0; i < a; i++ {
		sum += a
	}
	return sum
}

func main() {
	Add(1, 2)
	Sum(5)
	fmt.Println(Sum(5))
}
