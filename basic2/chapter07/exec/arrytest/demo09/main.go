package main

import "fmt"

func main() {
	var arr = [10]int{11, 26, 13, 62, 47, 59, 52, 26, 3, 32}
	sum := 0.0
	avg := 0.0

	for _, v := range arr {
		sum += float64(v)
	}
	avg = sum / float64(len(arr))

	fmt.Println(sum, avg)

	uppercount := 0
	lowercount := 0

	for _, v := range arr {
		if float64(v) > avg {
			uppercount++
		} else {
			lowercount++
		}
	}

	fmt.Println(uppercount, lowercount)

}
