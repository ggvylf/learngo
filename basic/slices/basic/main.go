package main

import "fmt"

func main() {

	//定义切片，切片是引用类型
	var s3 []int
	fmt.Printf("s3=%d,len=%d,cap=%d\n", s3, len(s3), cap(s3))

	//初始化
	s3 = []int{1, 2}
	fmt.Printf("s3=%d,len=%d,cap=%d\n", s3, len(s3), cap(s3))

	//声明并初始化，指定容量cap和长度len make(type,len,cap)
	//make()会在内存中分配空间
	s4 := make([]int, 3, 8)
	fmt.Printf("s4=%d,len=%d,cap=%d\n", s4, len(s4), cap(s4))

	//从数组得到切片，注意左包含右不包含  [start:end)

	arr := [...]int{1, 2, 3, 4, 5}
	//对arr切片
	s1 := arr[:]
	fmt.Println("s=", s1)
	s2 := arr[1:3]
	fmt.Println("s2=", s2)

	//遍历切片
	//方法1
	for i := 0; i < len(s1); i++ {
		fmt.Printf("s1[%d]=%v\n", i, s1[i])
	}

	//方法2
	for k, v := range s1 {
		fmt.Printf("s1[%d]=%v\n", k, v)
	}

	//空切片
	var sli1 []int
	sli2 := []int{}
	sli3 := make([]int, 0)

	fmt.Println(len(sli2), len(sli3))
	//判断空切片，判断len(slice)==0

	if len(sli1) == 0 {
		fmt.Println("sli1 is empty slice")
	}

}
