package main

import "fmt"

//接口是一个抽象类型，是一组method的集合，只关心对应的方法，不关心方法对应的属性
//例如计算各种多边形的面积。这里不关心多边形的类型，只关心计算面积的方法
//把具体对象有共同性的方法进行抽象，定义一个抽象类型，这就是接口类型。
//某个对象实现了这个接口类型中包含的方法，就可以认定为该对象就是对应的接口类型
//接口=接口类型+接口的值，类型和值都是动态的，可以存放任意的类型和和该类型对应的值

// 定义一个名字为speaker的接口类型，包含一个Speak()的方法
// 通常在对应的方法后边加er来命名
type Speaker interface {
	//定义一个发出叫声的方法
	Speak()
}

// 定义2个不同的结构体

// 定义猫
type Cat struct{}

// 定义猫叫的方法
func (c *Cat) Speak() {
	fmt.Println("喵喵喵")
}

// 定义狗
type Dog struct{}

// 定义狗叫的方法
func (d *Dog) Speak() {
	fmt.Println("汪汪汪")
}

func main() {

	// 创建一个Cat的实例
	c := &Cat{}
	c.Speak()

	// 创建一个Dog的实例
	d := &Dog{}
	d.Speak()

}
