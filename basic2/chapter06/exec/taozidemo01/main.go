package main

import "fmt"

func chi(n int) int {
	if n == 10 {
		return 1
	} else {
		return (chi(n+1) + 1) * 2
	}

}

func main() {
	day := 1

	fmt.Println(chi(day))
}
