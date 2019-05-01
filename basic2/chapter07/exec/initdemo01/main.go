package main

import "fmt"

func modify(arr *[3]int) {

	fmt.Printf("arr is %p,arr address is %p,arr value is %v\n", arr, &arr, *arr)
	arr[0] = 4
}

func main() {
	Arr01 := [3]int{1, 2, 3}
	fmt.Println("before=", Arr01)
	fmt.Printf("Arr01 add is %p\n", &Arr01)

	modify(&Arr01)

	fmt.Println("after=", Arr01)
}
