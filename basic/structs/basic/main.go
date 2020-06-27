package main

import "fmt"

//定义结构体
//值类型
type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {

	//声明结构体
	var stu1 Student

	//声明并初始化
	//返回的是指针
	var stu2 Student = Student{"lisi", 17, 88.8}

	//声明并初始化
	//返回的是指针
	var stu3 *Student = new(Student)

	//结构体字段赋值
	stu1.Name = "zhangsan"
	stu1.Age = 18
	stu1.Score = 99.8

	stu3 = &Student{
		Name:  "wangwu",
		Score: 88,
	}

	fmt.Println("stu1=", stu1)
	fmt.Println("stu2=", stu2)
	fmt.Println("stu3=", stu3)

	// 匿名结构体
	var s struct {
		name string
		age  int
	}
	s.name = "hehe"
	s.age = 8
	fmt.Println("s=", s)

}
