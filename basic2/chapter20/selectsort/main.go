package main

import "fmt"

func Sort(arr *[5]int) {

	for i := 0; i < len(arr)-1; i++ {
		//假设arr[i]是最大的
		max := arr[i]
		maxIndex := i

		//跳过arr[i]，遍历其他，已经排序过的不需要再比较
		for j := i + 1; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
				maxIndex = j
			}

		}

		//交换
		if maxIndex != 0 {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}

		fmt.Printf("第%d次,max=%d,arr=%v\n", i+1, max, arr)

	}

}

func main() {
	arr := [5]int{10, 34, 19, 100, 80}

	fmt.Println("before=", arr)
	Sort(&arr)
	fmt.Println("after=", arr)
}
