package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3, 4, 5}
	s1 := a[:]
	fmt.Println(s1, len(s1), cap(s1))

	for i := 0; i < len(s1); i++ {
		fmt.Printf("index=%d,value=%v\n", i, s1[i])
	}

	s2 := make([]int, 3, 6)

	for k, v := range s2 {
		fmt.Printf("index=%d,value=%v\n", k, v)
	}

	s3 := append(s2, 5)
	fmt.Println(s3)
	fmt.Println(s3, len(s3), cap(s3))
}
