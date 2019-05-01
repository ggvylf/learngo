package main

import "fmt"

func main() {
	var arr = [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	max := arr[0]
	min := arr[0]

	for _, v := range arr {
		if v > max {
			max = v

		}

	}

	for _, v := range arr {
		if v < min {
			min = v
		}

	}
	fmt.Println(max, min)

	sum := 0.0
	for _, v := range arr {
		sum += float64(v)
	}
	avg := (sum - float64(max) - float64(min)) / 6

	fmt.Println(avg)

	var b [8]float64

	for i := 0; i < len(arr); i++ {
		if float64(arr[i]) > avg {
			b[i] = float64(arr[i]) - avg
		} else {
			b[i] = avg - float64(arr[i])
		}

	}
	fmt.Println(b)

	good := b[0]
	bad := b[0]
	index := 0

	for k, v := range b {
		if float64(v) > bad {
			bad = float64(v)
			index = k
		}

	}
	fmt.Println(arr[index], index)

	for k, v := range b {
		if float64(v) < good {
			good = float64(v)
			index = k
		}

	}
	fmt.Println(arr[index], index)

}
