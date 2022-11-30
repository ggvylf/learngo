package main

import "fmt"

// go的继承
// 本质上是通过struct的嵌套来实现的 实际上是一种扩展
// go不支持重载

// 定义一个Animal类型的struct
type Animal struct {
	name string
}

// 定义Animal结构体的Move的方法
func (a Animal) Move() {
	fmt.Printf("%s is moving\n", a.name)
}

// 定义一个Dog类型的struct，扩展了Animal的struct
type Dog struct {
	Animal
	legs int
}

func (d Dog) Wang() {
	fmt.Printf("%s is wang\n", d.name)
}

// 重写了Move方法
func (d Dog) Move() {
	fmt.Printf("%s is use %dleg to moving\n", d.name, d.legs)
}

func main() {

	d1 := Dog{
		Animal: Animal{name: "dog1"},
		legs:   4,
	}

	fmt.Println(d1)

	// 明确调用Animal的Move方法
	d1.Animal.Move()

	// 使用的是Dog自己的Move方法
	d1.Move()

	d1.Wang()

}
