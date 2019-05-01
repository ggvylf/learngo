package main

import "fmt"

func main() {
	c1 := 'a'
	c2 := '0'
	fmt.Println(string(c1), c2)
	fmt.Printf("%c\n", c2)

	c3 := '世'
	fmt.Println(c3)
	fmt.Printf("%c\n%d\n", c3, c3)

	var c4 int = '界'
	fmt.Println(c4)
	fmt.Printf("%c\n%d\n", c4, c4)

	fmt.Println(string(c3 + c1))

}
