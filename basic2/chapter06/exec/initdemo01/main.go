package main

import "fmt"

var a = test()

func test() int {
	fmt.Println("test()")
	return 90
}

func init() {
	fmt.Println("this is init")

}

func main() {
	fmt.Println("this is main")
}
