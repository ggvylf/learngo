package main

import "fmt"

func fab(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fab(n-1) + fab(n-2)
	}
}

func main() {
	var n int
	n = 6
	result := fab(n)

	fmt.Println(result)
}
