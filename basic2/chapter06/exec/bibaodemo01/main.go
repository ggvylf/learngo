package main

import "fmt"

func AddUpper() func(int) int {
	n := 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))

}
