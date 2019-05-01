package main

import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex, findVal int) {

	if leftIndex > rightIndex {
		fmt.Println("not find,exit")
		return
	}

	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("find,index=%v\n", middle)
	}
}

func main() {

	arr := [6]int{1, 8, 10, 22, 89, 1000}
	BinaryFind(&arr, 0, len(arr)-1, 22)
}
