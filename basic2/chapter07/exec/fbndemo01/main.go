package main

import "fmt"

func fbn(n int) []uint64 {
	var arr = make([]uint64, n)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr

}

func main() {
	s := fbn(10)
	fmt.Println(s)

}
