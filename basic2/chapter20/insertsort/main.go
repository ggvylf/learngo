package main

import "fmt"

func Sort(arr *[5]int) {

	//假设arr[0]是一个有序数组，其他元素都是无序，插入到该有序数组中

	for i := 1; i < len(arr); i++ {

		//第一个要插入元素是arr[1]
		insertVal := arr[i]
		//插入的索引在元素的前一位
		insertIndex := i - 1

		//找到插入位置
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			//满足条件，原来元素向后移动
			arr[insertIndex+1] = arr[insertIndex]
			//插入的索引向前移动
			insertIndex--
		}

		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}

		fmt.Printf("第%d次，insertVal=%d，insertIndex+1=%d，arr=%v\n", i, insertVal, insertIndex+1, arr)

	}

}

func main() {
	arr := [5]int{23, 0, 12, 56, 34}

	fmt.Println("before=", arr)
	Sort(&arr)
	fmt.Println("after=", arr)
}
