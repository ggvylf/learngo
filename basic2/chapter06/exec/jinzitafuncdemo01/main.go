package main

import "fmt"

func jzt(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			if k == 1 || k == i*2-1 || i == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}

func main() {
	var a int
	fmt.Scanln(&a)
	jzt(a)

}
