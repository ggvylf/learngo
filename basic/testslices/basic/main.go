package main

import "fmt"

func main() {
	//声明arr
	arr := [...]int{1, 2, 3, 4, 5}

	//对arr切片
	s1 := arr[:]
	fmt.Println("s=", s1)

	s2 := arr[1:3]
	fmt.Println("s2=", s2)

	//创建切片
	var s3 []int
	fmt.Printf("s3=%d,len=%d,cap=%d\n", s3, len(s3), cap(s3))

	s4 := make([]int, 3, 8)
	fmt.Printf("s4=%d,len=%d,cap=%d\n", s4, len(s4), cap(s4))

	//遍历
	for i:=0;i<len(s1);i++ {
		fmt.Printf("s1[%d]=%v\n",i,s1[i])
	}

	for k,v:=range s1 {
		fmt.Printf("s1[%d]=%v\n",k,v)
	}
}
