package main

import "fmt"

func main() {

	arr := [3][]string{
		[]string{"1", "2", "3"},
		[]string{"4", "5", "6"},
		[]string{"7", "8", "9"},
	}

	fmt.Printf("origin arr=%v\n", arr)

	modify1(arr)
	fmt.Printf("after replace arr[2] arr=%v\n", arr)

	modify2(arr)
	fmt.Printf("after replace arr[2][0] arr=%v\n", arr)

}

func modify1(arr [3][]string) {
	arr[2] = []string{"a", "b", "c"}

}

func modify2(arr [3][]string) {
	arr[2][0] = "a"
}
