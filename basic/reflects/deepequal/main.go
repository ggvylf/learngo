package main

import (
	"fmt"
	"reflect"
)

// 比较2个map
func mapDeepEqual() {
	a := map[int]string{1: "a", 2: "b", 3: "c"}
	b := map[int]string{1: "a", 2: "b", 3: "c"}
	c := map[int]string{1: "a", 2: "b", 4: "c"}

	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(a, c))
}

// 比较切片
func sliceDeepEqual() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}

	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(a, c))
}

func main() {

	mapDeepEqual()

	sliceDeepEqual()

}
