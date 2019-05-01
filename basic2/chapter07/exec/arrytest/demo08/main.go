package main

import "fmt"

func main() {
	var arr = [10]int{90, 26, 78, 62, 47, 59, 52, 26, 91, 32}
	max := arr[0]
	min := arr[0]
	index := 0

	for k, v := range arr {
		if v > max {
			max = v
			index = k
		}

	}
	fmt.Println(max, index)

	for k, v := range arr {
		if v < min {
			min = v
			index = k
		}

	}
	fmt.Println(min, index)
}
