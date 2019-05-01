package main

import "fmt"

func main() {
	var arr = [4]int{1, 3, 5, 7}
	b := 6

	leftIndex := 0
	rightIndex := 0
	for i := 0; i < len(arr); i++ {
		if b > arr[i] {
			leftIndex = i
		} else {
			rightIndex = i
		}
	}
	fmt.Println(leftIndex, rightIndex)
	var arr2 [5]int
	for i := 0; i <= leftIndex; i++ {
		arr2[i] = arr[i]
	}
	arr2[rightIndex] = b
	for i := rightIndex + 1; i < len(arr2); i++ {
		arr2[i] = arr[i-1]
	}

	fmt.Println(arr2)

}
