package main

import "fmt"

// go的多态
// 多态就是用interface类型实现 只要实现了接口类型的方法 不同的对象就都可以使用同一个名字相同的方法

type Eater interface {
	Eat()
}

type Person struct {
	Name string
}
type Dog struct {
	Name string
}

func (p *Person) Eat() {
	fmt.Printf("%s is a person,eating\n", p.Name)
}

func (d *Dog) Eat() {
	fmt.Printf("%s is a dog,eating\n", d.Name)
}

func main() {
	p := &Person{
		Name: "tom",
	}
	d := &Dog{
		Name: "jerry",
	}

	p.Eat()
	d.Eat()

}
