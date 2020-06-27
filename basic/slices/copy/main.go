package main

import (
	"fmt"
)

func main() {

	//切片的赋值
	s1 := []int{1, 2, 3, 4, 5}

	s2 := s1

	fmt.Printf("before edit, s1=%v,s2=%v,s3=%v\n", s1, s2)

	//修改s2的值，等同于修改了底层数组的值，s1对应的值也会跟着改变
	s2[1] = 9
	fmt.Printf("after edit s2[1], s1=%v,s2=%v,s3=%v\n", s1, s2)

	//切片的copy
	c1 := []int{1, 2, 3}

	//c2必须初始化，并且len(c2)>=len(c1)，如果不初始化无法分配内存空间，也就无法完成copy()
	c2 := make([]int, len(c1))

	//拷贝c1的值到c2，c1和c2使用不同的内存空间
	copy(c2, c1)
	fmt.Println(c1, c2)
	fmt.Println("c2=", c2)

	//此时修改c1的值，不会影响c2
	c1[0] = 100
	fmt.Printf("c1=%v,c2=%v\n", c1, c2)

}
