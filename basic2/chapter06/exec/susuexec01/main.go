package main

import (
	"fmt"
	"math"
)

func isPrim(n float64) bool {
	s := int(math.Sqrt(n))
	for i := 2; i <= s; i++ {
		if int(n)%i == 0 {
			return false
		}

	}
	return true
}

func main() {
	for i := 101; i <= 200; i++ {

		if isPrim(float64(i)) {
			fmt.Println(i)
		}

	}

}
