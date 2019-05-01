package main

import (
	"fmt"
)

func main() {
	// var (
	// 	input int
	// 	n     int
	// )
	// n = rand.Intn(100)

	// for {
	// 	fmt.Scanf("%d\n", &input)
	// 	flag := false
	// 	switch {
	// 	case input == n:
	// 		fmt.Println("match")
	// 		flag = true

	// 	case input > n:
	// 		fmt.Println("bigger")

	// 	case input < n:
	// 		fmt.Println("smaller")
	// 	default:
	// 		fmt.Println("not null")
	// 	}

	// 	if flag {
	// 		break
	// 	}

	// }
	// fmt.Printf("the num is %d\n", n)

	str := "A"
	for i := 1; i < 6; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(str)
		}
		fmt.Println()

	}
}
