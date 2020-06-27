package main

import "fmt"

type Cat struct{}

type Animal struct{}

//值接受者实现接口和指针接收者实现接口的区别
//值接收者实现接口=值的方法+指针的方法
//指针接收者实现接口=指针的方法





//值接收者实现的方法
func (c Cat) Move() {
	fmt.Println("cat move")
}

func (c Cat) Eat() {
	fmt.Println("cat eat")
}

//指针接收者实现的方法
// func (c *Cat) Move() {
// 	fmt.Println("cat move")
// }

// func (c *Cat) Eat() {
// 	fmt.Println("cat eat")

//定义接口类型
type Animailer interface {
	Move()
	Eat()
}

func (a Animal) Alive() {
	Move(a)
	Eat(a)
}

func main() {

	c1 := Cat{}
	Alive(c1)

}
