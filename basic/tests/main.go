package main

import "fmt"

type Coder struct{}

func (c *Coder) Work() {
	fmt.Println("is working")
}

func (c *Coder) Rest() {
	fmt.Println("is resting")
}
func main() {

	var worker Worker
	c := new(Coder)
	worker = c

	switch worker.(type) {
	case Coder:
		fmt.Println("type is *Coder")
	case *Coder:
		fmt.Println("type is Coder")
	}

}
