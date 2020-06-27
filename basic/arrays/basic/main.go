package main

import "fmt"

func main() {
	//定义数组
	//必须制定数组的长度和类型，长度和类型是数组类型的一部分
	var a1 [3]int
	var a2 [2]bool
	var a4 [2]int

	//a1和a4的类型不同
	fmt.Printf("a1 type=%T,a2 type=%T\n", a1, a4)

	//不初始化，默认是0值
	//bool为false int和float都是0，string是""
	fmt.Printf("a1=%v,a2=%v\n", a1, a2)

	//初始化
	//方式1 先声明后赋值
	a1 = [3]int{1, 2, 3}
	a2 = [2]bool{true, false}
	fmt.Printf("a1=%v,a2=%v\n", a1, a2)

	//方式2 声明同时赋值
	a3 := [3]string{"hello", "world", "go"}
	fmt.Println("a3=", a3)

	//长度推断
	a5 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("arr5=%v,arr5 len=%v\n", a5, len(a5))

	//方式3 索引指定
	a6 := [5]int{0: 1, 3: 4}
	fmt.Println("a6=", a6)

	//数组的容量和长度
	fmt.Println("a1 len =", len(a1))
	fmt.Println("a1 cap =", cap(a1))

	//遍历数组
	//方式1
	var arr1 [4]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
		fmt.Printf("arr[%d]=%v\n", i, arr1[i])
	}

	//方式2
	for index, value := range arr1 {
		fmt.Printf("arr1[%d]=%v\n", index, value)
	}

	//多维数组
	//方式1
	m1 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("m1=", m1)

	m2 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}

	fmt.Println("m2=", m2)

	//多维数组遍历
	for _, v1 := range m1 {
		fmt.Println("v1=", v1)
		for _, v2 := range v1 {
			fmt.Println("v2=", v2)
		}
	}

	//数组是值类型，拷贝也是值拷贝，修改拷贝不会影响原始数组
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 4
	fmt.Printf("b1=%v,b2=%v\n", b1, b2)
	b1[0] = 3
	fmt.Printf("b1=%v,b2=%v\n", b1, b2)
}
