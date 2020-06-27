package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) Move() {
	fmt.Printf("%s is moving\n", a.name)
}

type Dog struct {
	Animal
	legs int
}

func (d Dog) Wang() {
	fmt.Printf("%s is wangwangwang\n", d.name)
}

func main() {

	d1 := Dog{
		Animal: Animal{name: "dog1"},
		legs:   4,
	}

	fmt.Println(d1)

	d1.Move()
	d1.Wang()

}
