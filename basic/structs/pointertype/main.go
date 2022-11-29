package main

import (
	"fmt"
)

type Person struct {
	Name, Gender string
}

// 函数内部操作的修改的是传入变量的值拷贝
func f1(p Person) {
	p.Gender = "female"

}

// 传递指针可以直接修改传入变量的值
func f2(p *Person) {
	//语法糖
	//标准写法 (*p).Gender="female"
	p.Gender = "female"

}

func main() {
	p := Person{"tom", "male"}

	f1(p)
	fmt.Println("p=", p)

	//直接传入结构体的指针
	f2(&p)
	fmt.Println("p=", p)

	//使用new()实例化struct，返回的是struct的直针
	p2 := new(Person)
	fmt.Printf("p2 type=%T\n", p2)
	p2.Name = "jerry"
	p2.Gender = "male"
	f2(p2)

	fmt.Println("p2=", p2)

}
