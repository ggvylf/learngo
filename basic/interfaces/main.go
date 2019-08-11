package main

import "fmt"

//结构体
type Person struct {
	Name string
	Age  int
}
type Studnent struct {
	Person
	Score int
}

//定义一个live接口，只要能实现live内的所有Method，就认为该对象是live类型的
type live interface {
	Eat()
	Sleep()
}

//具体对象来实现方法，注意这是method，不是function。
func (p *Studnent) Sleep() {
	fmt.Printf("%s is sleeping\n", p.Name)
}

func (p *Studnent) Eat() {
	fmt.Printf("my name is %s,my age is %d,myscore is %d,im is eating\n", p.Name, p.Age, p.Score)
}

func (p *Person) Sleep() {
	fmt.Printf("%s is sleeping\n", p.Name)
}

func (p *Person) Eat() {
	fmt.Printf("my name is %s,my age is %d, im is eating\n", p.Name, p.Age)
}

//定义了一个函数，接受参数是live类型。
func life(l live) {
	l.Eat()
	l.Sleep()
}

func main() {
	p1 := &Person{
		Name: "tom",
		Age:  18,
	}

	s1 := &Studnent{}
	s1.Name = p1.Name
	s1.Age = p1.Age
	s1.Score = 13

	life(p1)

	life(s1)

	//声明了一个live类型
	var l live
	s2 := &Studnent{}
	s2.Name = "jerry"
	//s2因为实现了live的所有方法，所以可以看成live类型，可以直接赋值给l。
	l = s2
	fmt.Println(l)

}
