package main

import "fmt"

//结构体的声明
type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {

	//结构体初始化的几种方式
	var stu Student
	//返回的是指针
	var stu2 *Student = &Student{"lisi", 17, 88.8}
	//返回的是指针
	var stu3 *Student = new(Student)

	stu.Name = "zhangsan"
	stu.Age = 18
	stu.Score = 99.8

	stu3.Name = "wangwu"
	stu3.Age = 16
	stu3.Score = 77.7

	fmt.Println(stu)
	fmt.Println(stu2)
	fmt.Println(stu3)
}
