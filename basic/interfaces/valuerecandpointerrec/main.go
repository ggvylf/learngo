package main

import "fmt"

type Cat struct{}

//值接收者实现接口和指针接收者实现接口的区别
//值接收者实现接口=值的方法+指针的方法+值的结构体+指针的结构体
//指针接收者实现接口=指针的方法+指针的结构体

//值接收者实现的方法
// func (c Cat) Move() {
// 	fmt.Println("cat move")
// }

// func (c Cat) Eat() {
// 	fmt.Println("cat move")
// }

//指针接收者实现的方法
func (c *Cat) Move() {
	fmt.Println("cat move")
}

func (c *Cat) Eat() {
	fmt.Println("cat move")
}

//定义接口类型
type Animal interface {
	Move()
	Eat()
}

func main() {

	c1 := Cat{}
	c2 := &Cat{}

	var a1 Animal

	//指针接收者的方法不能接收值类型的，但是值接收者的方法可以同时接受值和指针类型
	//a1=&c1就可以了
	a1 = c1
	a1 = c2
	fmt.Println(a1)

}
