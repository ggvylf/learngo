package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Studnent struct {
	Person
	Score int
}

type Test interface {
	Print()
	Sleep()
}

func (p *Studnent) Sleep() {
}

func (p *Studnent) Print() {
	fmt.Printf("%s\n%d\n%d\n", p.Name, p.Age, p.Score)
}

func (p *Person) Sleep() {
}

func (p *Person) Print() {
	fmt.Printf("%s\n%d\n", p.Name, p.Age)
}

func main() {
	var t Test
	stu := &Studnent{}
	stu.Name = "tom"
	stu.Age = 19
	stu.Score = 88

	t = stu
	t.Print()
	t.Sleep()

	p := &Person{}
	p.Name = "jerry"
	p.Age = 15
	t = p
	t.Print()

	//空接口
	var a int
	var b interface{}
	b = a
	fmt.Printf("%T\n", b)
}
