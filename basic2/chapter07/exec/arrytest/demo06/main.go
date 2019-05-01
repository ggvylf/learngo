package main

import "fmt"

func main() {
	arr := [10]string{"a", "b", "AA", "c", "d", "AA", "cdsdf", "asdf", "abc", "asdf"}
	flag := false
	for k, v := range arr {
		if v == "AA" {
			fmt.Println("index=", k)
			flag = true

		} else if k == len(arr)-1 && flag == false {
			fmt.Println("not found")

		}
	}
}
