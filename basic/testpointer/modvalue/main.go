package main

import "fmt"

func main() {

	a := 5
	fmt.Println("before mod,a=", a)

	p := &a

	*p = 3

	fmt.Println("after mod,a=", a)

}
