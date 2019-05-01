package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// func (stu *Student) String() string {
// 	str := fmt.Sprintf("name=[%v],age=[%v]\n", stu.Name, stu.Age)
// 	return str
// }

func main() {
	stu := Student{
		"tom",
		18,
	}

	fmt.Println(&stu)
}
