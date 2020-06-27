package main

import "fmt"

//定义猫和狗
type Cat struct{}
type Dog struct{}

//定义叫的函数

func (c Cat) Speak() {
	fmt.Println("喵喵喵")
}

func (d Dog) Speak() {
	fmt.Println("汪汪汪")
}



//接口是一个抽象类型，是一组method的集合，只关心对应的方法，不关心方法对应的属性
//例如计算各种多边形的面积。这里不关心多边形的类型，只关心计算面积的方法
//把具体对象有共同性的方法进行抽象，定义一个抽象类型，这就是接口类型。
//某个对象实现了这个接口类型中包含的方法，就可以认定为该对象就是对应的接口类型


//接口=接口类型+接口的值，类型和值都是动态的，可以存放任意的类型和和该类型对应的值


//定义一个名字为speaker的接口，通常在对应的方法后边加er来命名
type speaker interface {
	//所有实现了Speak()方法的类型，都可以看做是speaker类型
	Speak()
}



//定义一个打的函数，接收一个接口类型的speaker参数，然后调用对应的Speak()
func Da(x speaker) {
	x.Speak()
}






func main() {
	var c1 Cat
	var d1 Dog

	Da(c1)
	Da(d1)


	//spk是一个接口类型
var spk speaker
//实现了该接口变量都可以任意赋值，可以看做都是speaker类型的
spk=c1
fmt.Printf("%T\n",spk)
}
