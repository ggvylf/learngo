package main

import (
	"encoding/json"
	"fmt"
)

// 指定了struct字段的tag，格式为kv对,多个tag用空格分割
// 内部是通过反射reflect来实现的
// 在高qps场景下 reflect的效率并不高

type Person struct {
	Name string `json:"my_name"  db:"name"  xml:”name“`
	age  int    `json:"my_age"`
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		age:  age,
	}
}

// 使用自定义方法来返回私有变量的值拷贝
func (p *Person) GetAge() int {
	return p.age
}

type Class struct {
	Title string   `json:"my_title"`
	Stu   []Person `json:"my_stu_list"`
}

func main() {
	c1 := Class{
		Title: "my class",
		Stu:   make([]Person, 0, 20),
	}

	for i := 0; i < 10; i++ {
		tmpStu := NewPerson(fmt.Sprintf("stu%02d", i), i)
		c1.Stu = append(c1.Stu, *tmpStu)
	}

	//序列化
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("c1 json.marshal=", string(data))
	fmt.Println()

	//反序列化
	//定义一个结构和Class结构体一样的interface，用来存放unmarshal的数据。
	var c2 Class
	err = json.Unmarshal(data, &c2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("c2 json.unmarshl=", c2)

}
