package main

import (
	"fmt"
	"math"
	"time"
)

func GetSqrt(i int, chanSema chan int) {
	chanSema <- i
	result := math.Sqrt(float64(i))
	fmt.Printf("%d的平方根＝%.2f\n", i, result)
	time.Sleep(time.Second)
	<-chanSema

}

func main() {

	chanSema := make(chan int, 5)

	for i := 1; i <= 100; i++ {
		go GetSqrt(i, chanSema)
	}

	time.Sleep(26 * time.Second)
	fmt.Println("main func end")

}
