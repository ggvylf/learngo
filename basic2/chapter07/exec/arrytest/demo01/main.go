package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a [10]int
	var b [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)

	//倒序显示
	for i := 0; i < len(a); i++ {
		b[i] = a[len(a)-i-1]
	}
	fmt.Println(b)

	//平均值
	sum := 0.0
	for _, v := range a {
		sum += float64(v)
	}
	fmt.Println("avg=", sum/float64(len(a)))

	//最大值和最大值的下标
	max := a[0]
	index := 0
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
			index = i
		}

	}
	fmt.Printf("最大值为%d,索引为%d\n", max, index)

	//查找里边是否有55
	for k, v := range a {
		if v == 55 {
			fmt.Println("找到了")
			break
		} else if k == len(a)-1 {
			fmt.Println("没找到")
		}
	}

}
