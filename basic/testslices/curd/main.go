package main

import "fmt"

func main() {

	//追加元素
	var s1 []int
	s1 = append(s1, 1)
	s1 = append(s1, 2, 3, 4)
	fmt.Println(s1)

	//追加slice，注意slice后要加...
	s2 := []int{5, 6}
	s1 = append(s1, s2...)
	fmt.Println(s1)

	//
}
