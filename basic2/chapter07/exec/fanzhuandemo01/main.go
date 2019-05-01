package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var a = [5]int{}
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	temp := 0
	for i := 0; i < len(a)/2; i++ {
		temp = a[i]
		a[i] = a[len(a)-i-1]
		a[len(a)-i-1] = temp

	}
	fmt.Println(a)
}
