package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

//值接收者，传入的是值拷贝，通常用在不需要修改结构体数据的场景下
func (p Person) AddAge1() {
	p.Age++
}

//指针接收者，可以修改结构体中的数据
func (p *Person) AddAge2() {
	p.Age++
}

func main() {
	p1 := NewPerson("tom", 18)

	//值接收者，不会修改p1的值
	p1.AddAge1()
	fmt.Println("p1=", p1)

	//p1.age=19
	p1.AddAge2()
	fmt.Println("p1=", p1)

	// p2 := Person{"jerry", 18}

}
