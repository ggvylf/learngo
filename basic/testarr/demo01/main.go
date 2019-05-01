package main

import "fmt"

func main() {
	//声明后赋值
	var arr2 [3]string
	arr2[0] = "tom"
	arr2[2] = "jerry"

	fmt.Printf("arr2=%v\n", arr2)

	//声明的同时赋值
	var arr3 [3]int = [3]int{4, 5, 6}
	fmt.Println("arr3=", arr3)

	arr4 := [3]int{7, 8, 9}
	fmt.Println("arr4=", arr4)

	//长度推断
	arr5 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr5=", arr5)
	fmt.Println("arr5 len=", len(arr5))

	//遍历数组
	var arr1 [4]int
	for i := 0; i < 4; i++ {
		arr1[i] = i
		fmt.Printf("arr[%d]=%v\n", i, arr1[i])
	}

	for index, value := range arr1 {
		fmt.Printf("arr1[%d]=%v\n", index, value)
	}

	//数组的容量和长度
	fmt.Println("arr len =", len(arr1))
	fmt.Println("arr cap =", cap(arr1))

}
