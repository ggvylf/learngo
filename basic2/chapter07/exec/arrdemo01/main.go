package main

import "fmt"

func main() {
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("input the %d num:", i+1)
		fmt.Scanln(&score[i])
	}

	for k, v := range score {
		fmt.Printf("index=%d,vaule=%v\n", k, v)
	}
}
