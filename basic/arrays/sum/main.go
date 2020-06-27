package main

import (
	"fmt"
)

func main() {

	a1 := [5]int{1, 2, 3, 4, 5}
	sum := 0

	//求数组总和
	for _, v := range a1 {
		sum += v
	}

	fmt.Println("sum=", sum)

	//打印出数组总2个元素之和=6的下标
	for key1, _ := range a1 {
		for key2, _ := range a1 {
			if a1[key1]+a1[key2] == 6 && key1 < key2 {
				fmt.Printf("key1=%v,key2=%v\n", key1, key2)
			}
		}
	}

	//外层循环从索引0开始
	//内层循环从外层索引开始
	for i := 0; i < len(a1); i++ {
		for j := i; j < len(a1); j++ {
			if a1[i]+a1[j] == 6 && i != j {
				fmt.Printf("i=%v,j=%v\n", i, j)
			}
		}

	}
}
