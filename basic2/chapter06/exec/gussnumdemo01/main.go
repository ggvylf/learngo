package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 10
	i := 1

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	fmt.Println(n)

	for i = 1; i <= count; i++ {
		var g int
		fmt.Print("input=")
		fmt.Scan(&g)

		if g == n {
			fmt.Println("bingo")
			break
		}
	}

	switch i {
	case 1:
		fmt.Println("天才")
	case 2, 3:
		fmt.Println("聪明")
	case 4, 5, 6, 7, 8, 9:
		fmt.Println("一般般")
	case 10:
		fmt.Println("可算猜对了")
	default:
		fmt.Println("一次都没猜中")
	}

}
