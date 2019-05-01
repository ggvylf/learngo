package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("start defer")
		if p := recover(); p != nil {
			fmt.Println("get painc")
		}
		fmt.Println("stop defer")
	}()
	fmt.Println("main start")

	panic(errors.New("painc"))
	fmt.Println("main stop")
}
