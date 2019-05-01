package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//值
	var p1 Person
	p1.Name = "tom"
	p1.Age = 18
	fmt.Println("p1=", p1)

	//有序，所有字段必须全赋值
	p2 := Person{"jerry", 18}
	fmt.Println("p2=", p2)
	//无序，只给制定字段赋值
	p3 := Person{Name: "mike"}
	fmt.Println("p3=", p3)

	//指针结构体
	// var p4 *Person = new(Person)
	p4 := new(Person)
	//标准写法
	(*p4).Name = "ann"
	//语法糖，可以不写(*p4).Age
	p4.Age = 30
	fmt.Println("p4=", p4)

	// var p5 *Person = &Person{"jim", 18}
	p5 := &Person{"jim", 18}
	fmt.Println("p5=", p5)

}
