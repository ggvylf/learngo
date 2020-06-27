package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 结构体的构造函数
// 结构体是值类型，通常建议返回指针，返回值在字段多的情况下会造成内存开销
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	p := NewPerson("tom", 18)
	fmt.Println("p=", p)

}
