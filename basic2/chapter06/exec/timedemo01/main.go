package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 1000; i++ {
		str += strconv.Itoa(i)
	}

}

func main() {
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Println("use time=", end-start)

}
