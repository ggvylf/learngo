package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BinaryFind(arr *[10]int, leftIndex int, rightIndex, findVal int) {

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
	rand.Seed(time.Now().UnixNano())
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	temp := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)

	BinaryFind(&arr, 0, len(arr)-1, 90)

}
