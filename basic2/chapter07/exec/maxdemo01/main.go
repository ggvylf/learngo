package main

import "fmt"

func main() {
	var a = [5]int{112, 223, 32, 444, 124}
	max := a[0]
	maxindex := 0
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
			maxindex = i
		}

	}
	fmt.Printf("max=%d,maxindex=%d\n", max, maxindex)
}
