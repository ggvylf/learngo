package main

import "fmt"

func main() {
	// 示例1。
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	s1 := a1[:]

	//直接切片
	s2 := s1[:3]

	//make()后copy()
	s3 := make([]int, 3)
	copy(s3, s1)
	fmt.Printf("s1= %v,len=%v,cap=%v,addr=%p\n", s1, len(s1), cap(s1), &s1)
	fmt.Printf("s2= %v,len=%v,cap=%v,addr=%p\n", s2, len(s2), cap(s2), &s2)
	fmt.Printf("s3= %v,len=%v,cap=%v,addr=%p\n", s3, len(s3), cap(s3), &s3)
}
