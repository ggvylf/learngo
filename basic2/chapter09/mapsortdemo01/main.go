package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := map[int]int{
		10: 100,
		1:  13,
		4:  11,
	}

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	//输出map1对应的value
	for _, v := range keys {
		fmt.Printf("map1[%v]=%v\n", v, map1[v])
	}

}
