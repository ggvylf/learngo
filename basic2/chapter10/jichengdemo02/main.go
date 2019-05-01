package main

import "fmt"

type Student struct {
	Name string
	age  int
}

type Puple struct {
	*Student
	Name string
}

func main() {
	p := &Puple{
		&Student{
			"tom", 13,
		},
		"jerry",
	}

	p2 := &Puple{
		&Student{"a", 13}, "b"}

	fmt.Println(p.Student.Name)

	fmt.Println(p2.Student.age)

}
