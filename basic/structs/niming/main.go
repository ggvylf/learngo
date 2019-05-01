package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//Student中也有相同的Age字段
type Student struct {
	Person
	Score int
	Age   int
}

func (stu *Student) init(name string, age int, score int) {

	stu.Name = name
	stu.Age = age
	stu.Score = score
	fmt.Println(*stu)

}

func main() {

	stu := Student{}
	stu.init("tom", 1, 3)

	fmt.Println(stu.Age)
	fmt.Println(stu)
}
