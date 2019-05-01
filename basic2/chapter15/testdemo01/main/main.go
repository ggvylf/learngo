package main

import "fmt"

func addUpper(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

func main() {

	result := addUpper(10)
	fmt.Println(result)

	if result == 55 {
		fmt.Println("ok")
	} else {
		fmt.Println("failed")
	}

}
