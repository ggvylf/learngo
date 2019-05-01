package main

import "fmt"

func addUpper(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

func Sub(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(addUpper(10))
	fmt.Println(Sub(1, 2))

}
