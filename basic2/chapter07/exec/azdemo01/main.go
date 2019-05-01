package main

import "fmt"

func main() {
	var a = [26]byte{}

	for i := 0; i < len(a); i++ {
		a[i] = byte('A' + i)
		fmt.Printf("index=%d;value=%c\n", i, a[i])
	}
}
