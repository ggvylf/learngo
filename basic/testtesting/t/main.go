package main

import "fmt"

func Print1(i int) int {
	return i + 1
}

func Print2(i int) int {
	return i + 2
}

func main() {
	fmt.Println("func Print1", Print1(2))
	fmt.Println("func Print2", Print2(2))

}
