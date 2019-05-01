package main

import "fmt"

func main() {
	arr := [5]int{33, 6, 24, 86, 55}

	//长度为n的数组，需要排序n-1次

	//i控制排序的次数
	for i := 1; i < len(arr); i++ {
		//j控制每次排序中需要操作的元素个数
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	fmt.Println(arr)
}
