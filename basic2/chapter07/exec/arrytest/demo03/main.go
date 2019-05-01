package main

import "fmt"

func main() {
	var arr [3][4]int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = 1
		}

		fmt.Println(arr[i])
	}

	fmt.Println()

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || (i == len(arr)-1) || j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 0
			} else {
				arr[i][j] = 1
			}

		}
		fmt.Println(arr[i])

	}

}
