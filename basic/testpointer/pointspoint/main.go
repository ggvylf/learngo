package main

import "fmt"

func main() {
	a := 5
	p := &a
	fmt.Printf("p type=%T,p=%v,*p=%v,&p=%v\n", p, p, *p, &p)

	p2 := &p
	fmt.Printf("p2 type=%T,p2=%v,*p2=%v,&p2=%v\n", p2, p2, *p2, &p2)
}
