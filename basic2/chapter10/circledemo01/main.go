package main

import (
	"fmt"
)

type Circle struct {
	redius float64
}

func (c *Circle) Area() float64 {
	sum := 3.14 * c.redius * c.redius
	return sum

}

func main() {
	var c Circle
	c.redius = 4.0
	fmt.Printf("%.2f\n", c.Area())

}
