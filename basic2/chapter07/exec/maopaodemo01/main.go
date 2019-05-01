package main

import "fmt"

func Bubble(arr *[5]int) {
	fmt.Println("before=", *arr)
	temp := 0
	//外层循环，循环数组长度-1
	for i := 0; i < len(arr)-1; i++ {
		//内层循环，相邻两数比较
		//循环次数递减,已排好序的元素不需要再次排序
		for j := 0; j < len(arr)-1-i; j++ {
			//如果需要从大到小排列，改为<即可
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}

	}
	fmt.Println("after=", (*arr))

}

func main() {
	arr := [5]int{22, 34, 62, 13, 7}

	Bubble(&arr)

}
