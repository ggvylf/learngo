package main

import "fmt"

//创建自定义类型
type MyInt int

//自定义类型的方法
func (m MyInt) Hello() {
	fmt.Println("hello world")
}
func main() {

	i := MyInt(3)

	fmt.Printf("i type=%T\n", i)

	//等同于给int类型增加了Hello()方法
	i.Hello()

}
