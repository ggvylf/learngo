package main

import "fmt"

func main() {
	var arr = [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	var tmp [4]int

	tmp = arr[0]
	arr[0] = arr[3]
	arr[3] = tmp

	tmp = arr[1]
	arr[1] = arr[2]
	arr[2] = tmp

	fmt.Println(arr)

	
}
