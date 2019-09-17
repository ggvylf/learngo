package main

import (
	"fmt"
	"math"
)

func area(r float64) float64 {
	if r <= 0 {
		panic("半径不能为负数")
	}
	return math.Pi * r * r
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic的原因=", err)
		}
	}()

	res := area(float64(-3))

	fmt.Println("面积=", res)
}
