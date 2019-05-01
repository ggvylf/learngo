package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())
	ch1 := time.After(5 * time.Second)

	time := <-ch1
	fmt.Println(time)
}
