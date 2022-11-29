package main

import "fmt"

// 定义结构体
// 值类型
type Person struct {
	Name string
	Age  int
}

// 结构体嵌套
type Student struct {
	//匿名嵌套，可以直接访问该结构体下的变量。如果多个匿名结构体有相同的字段，只能写全了用来区分
	//同时这也是继承
	Person
	Score float32
}

// 结构体的匿名字段，字段类型必须唯一
type Anno struct {
	string
	int
}

func main() {

	//结构体的初始化
	NewStruct()

	// 匿名结构体
	var s struct {
		name string
		age  int
	}
	s.name = "hehe"
	s.age = 8
	fmt.Println("s=", s)

	a1 := Anno{"aaa", 123}
	//匿名字段你的调用方法
	fmt.Println(a1.string, a1.int)
}

func NewStruct() {
	//创建结构体实例
	var stu1 Student

	//结构提字段复制
	stu1.Name = "zhangsan"
	stu1.Age = 18
	stu1.Score = 99.8

	//声明并初始化
	var stu2 Student = Student{Person{"lisi", 17}, 88.8}

	//使用new()初始化，返回的是struct的指针
	//等同于 stu3:=&Student{}
	// stu3类型是*Student{}
	stu3 := new(Student)

	//给结构体中的部分字段赋值
	//注意： 嵌套的结构体，自己的字段必须有值
	stu3 = &Student{Person{}, 88}
	fmt.Println("stu1=", stu1)
	fmt.Println("stu2=", stu2)
	fmt.Println("stu3=", stu3)

}
