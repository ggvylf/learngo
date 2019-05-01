package main

import "fmt"

func Sort(left, right int, arr *[7]int) {
	//2个索引标记
	arrleft := left
	arrright := right
	//中间点
	pivot := arr[(left+right)/2]

	//完成排序
	for left < right {
		//从左边找到比pivot大的元素，直到找到
		for arr[left] < pivot {
			//索引标记向右移动
			left++
		}
		//从右边找到比piovt小的元素,直到找到
		for arr[right] > pivot {
			//向左移动
			right--
		}

		//基线条件
		if left >= right {
			break
		}

		//交换找到的元素
		arr[left], arr[right] = arr[right], arr[left]

		//左边和中间点相等，right标记向前移动一个，跳出循环
		if arr[left] == pivot {
			right--
		}

		//右边和中间点相等，left前后移动一个，跳出循环
		if arr[right] == pivot {
			left++
		}
		fmt.Println(arr)

	}
	//相等就在移动一下
	if left == right {
		left++
		right--
	}
	//向左递归
	if arrleft < right {
		Sort(arrleft, right, arr)
	}

	//向右递归
	if arrright > left {
		Sort(left, arrright, arr)
	}

}

func main() {
	arr := [7]int{-9, 78, 0, 23, -567, 70, 3}

	fmt.Println("before=", arr)
	Sort(0, len(arr)-1, &arr)
	fmt.Println("after=", arr)
}
