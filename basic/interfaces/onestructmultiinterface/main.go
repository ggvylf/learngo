package main

import "fmt"

//同一个struct可以实现多个接口

type Cat struct{}

//定义了2个接口
type mover interface {
	Move()
}

type eater interface {
	Eat()
}

//实现了2个接口

func (c Cat) Move() {
	fmt.Println("a is moving")

}

func (C Cat) Eat() {
	fmt.Println("a is eating")
}

//接口的嵌套
type Animal interface {
	mover
	eater
}

func main() {
}
