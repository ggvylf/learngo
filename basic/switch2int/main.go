package main

import "fmt"

func swap(a, b *int) {
	fmt.Printf("the pointer address is &a=%p,&b=%p\n", &a, &b)
	fmt.Printf("the pointer's value is *a=%d,*b=%d\n", *a, *b)
	*a, *b = *b, *a
	fmt.Printf("swaped the value is *a=%d,*b=%d\n", *a, *b)
	return
}

func main() {
	a := 100
	b := 200
	fmt.Printf("the memory address is a=%p,b=%p\n", &a, &b)
	swap(&a, &b)
	fmt.Printf("after swaped value is a=%d,b=%d\n", a, b)
}
