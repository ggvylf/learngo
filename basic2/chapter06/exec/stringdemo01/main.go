package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world you and me"
	result := strings.Fields(str)
	fmt.Printf("type is %T,value is %v\n", result, result)
	for k, v := range result {
		fmt.Printf("inidex=%d,value=%v\n", k, v)
	}
}
