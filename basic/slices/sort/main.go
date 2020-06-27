//对slice排序

package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int = []int{1, 3, 5, 99, 8, 44}
	var b []string = []string{"b", "d", "z", "a", "aaadfb", "AAASD", "Z", "D"}
	var c []float64 = []float64{1.23, 4.56, 0.04, 0.003}
	//int 排序
	sort.Ints(a)
	fmt.Printf("int after sort:%v\n", a)

	//string 排序
	sort.Strings(b)
	fmt.Printf("string after sort:%v\n", b)

	//float排序
	sort.Float64s(c)
	fmt.Printf("float after sort:%v\n", c)

	//int查找，返回index
	index1 := sort.SearchInts(a, 3)
	index2 := sort.SearchInts(a, 77)
	fmt.Println(index1)
	fmt.Println(index2)
	fmt.Println(len(a))

	//string查找
	index3 := sort.SearchStrings(b, "Q")
	fmt.Println(b[index3])
	fmt.Println(len(b))

	//float查找
	index4 := sort.SearchFloat64s(c, 4.56)
	fmt.Println(index4)

}
