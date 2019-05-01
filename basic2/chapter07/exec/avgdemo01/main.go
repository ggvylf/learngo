package main

import "fmt"

func main() {
	sum := 0
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	for _, v := range a {
		sum += v
	}

	var avg float64
	avg = float64(sum) / float64(len(a))

	fmt.Printf("sum=%d,avg=%.2f\n", sum, avg)
}
