package main

import "fmt"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age`
}

func (person Person) speak() {
	fmt.Printf("%s is good\n", person.Name)
}

func (person Person) jisuan() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("1++1000=", sum)
}

func (person Person) jisuan2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1++%d=%d\n", n, sum)

}

func (person Person) getSum(a, b int) int {
	return a + b
}

func main() {
	p := Person{"tom", 18}

	p.speak()
	p.jisuan()
	p.jisuan2(10)
	fmt.Println("1+3=", p.getSum(1, 3))

}
