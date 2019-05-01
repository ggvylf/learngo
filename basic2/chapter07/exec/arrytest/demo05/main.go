package main

import "fmt"

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i*2 + 1
	}

	fmt.Println(arr)
	var tmp int
	for i := 0; i < len(arr)/2-1; i++ {
		tmp = arr[i]
		arr[i] = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = tmp
	}
	fmt.Println(arr)
}
