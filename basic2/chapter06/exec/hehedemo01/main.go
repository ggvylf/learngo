package main

import "fmt"

func main() {
	count := 0
	for i := 1; i < 100; i += 2 {
		fmt.Printf("%d ", i)
		count++
		if count%5 == 0 {
			fmt.Println()
			count = 0
		}
	}
}
