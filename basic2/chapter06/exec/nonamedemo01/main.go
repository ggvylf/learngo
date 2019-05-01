package main

import "fmt"

var Res = func(a, b int) int {
	return a + b
}

func main() {
	a := 2
	b := 3

	res2 := Res(a, b)
	fmt.Println(res2)
	res3 := Res(10, 20)
	fmt.Println(res3)
}
