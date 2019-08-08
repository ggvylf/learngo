package main

import "fmt"

//定义struct
type Person struct {
	Name string
	age  int
}

//工厂模式，构造函数
func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

//方法，用来获取struct中小写的变量
func (p *Person) GetAge() int {
	return p.age
}

func main() {
	p := NewPerson("tom", 13)
	fmt.Println(p.Name, p.GetAge())
}
