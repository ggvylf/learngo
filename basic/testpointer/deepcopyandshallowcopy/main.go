package main

import "fmt"

func main() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//深拷贝
	arr2:=arr1
	arr2[0]=5
	fmt.Println(arr1,arr2)


	//浅拷贝
	arr3 := &arr1
	(*arr3)[0] = 5
	arr3[2]=7
	fmt.Println(arr1,arr3)
}
