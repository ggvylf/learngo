package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("not convert to int")
	} else {
		fmt.Printf("convert ok, the value is %d\n", num)
	}
}
