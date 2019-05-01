package main

import "fmt"

func Swap1(a, b *int) {
	t := *b
	*b = *a
	*a = t

}

func Swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a := 3
	b := 5
	Swap1(&a, &b)
	fmt.Println(a, b)

	c := 7
	d := 8
	fmt.Println(Swap2(c, d))
}
