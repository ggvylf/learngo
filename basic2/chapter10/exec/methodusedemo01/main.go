package main

import "fmt"

type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	fmt.Print()

}

func (mu MethodUtils) Jx(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}

func (mu MethodUtils) Mj(len, width float64) float64 {
	return len * width
}

func main() {
	var mu MethodUtils

	mu.Print()
	fmt.Println()
	mu.Jx(2, 6)
	fmt.Println("矩形面积=", mu.Mj(2.5, 8.7))

}
