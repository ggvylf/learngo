package main

import "fmt"

func main() {
	var i int = 10
	fmt.Printf("i value is %d,i address is %v\n", i, &i)

	p := &i

	fmt.Printf("pointer type is %T\np value is %v\np memory address is %v\n*p = i,*p value is %v\n", p, p, &p, *p)

	*p = 100
	fmt.Println("i value changed, new value is ", i)

}
