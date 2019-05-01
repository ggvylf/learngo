package main

import "fmt"

func main() {
	arr := [5]string{"hello", "world", "go", "tom", "jerry"}

	str := "go"

	//方式1
	for k, v := range arr {
		if str == v {
			fmt.Println("found")
			fmt.Println(k, v)
			break
		} else if k == len(arr)-1 {
			fmt.Println("not found")
		}

	}

	//方式2
	index := -1
	for k, v := range arr {
		if str == v {
			index = k
			break
		}
	}
	if index != -1 {
		fmt.Println(index, arr[index])
	} else {
		fmt.Println("not found")
	}

}
