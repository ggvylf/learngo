//链表
package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	Next  *Student
}

func list(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next

	}
}

//尾部插入
func insertTail(p *Student) {
	var tail = p
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}

		tail.Next = &stu
		tail = &stu
	}
	list(tail)

}

func insertHead(p *Student) {
	var head = p
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}

		stu.Next = head
		head = &stu
		head.Next = stu.Next

	}
	list(head)

}

func delNode(p *Student) {
	var priv *Student = p
	for p != nil {
		if p.Name == "stu6" {
			priv.Next = p.Next
			break
		}
		priv = p
		p = p.Next

	}
	list(priv)
}

func main() {
	//无序初始化并返回指针
	head := &Student{
		Name:  "zhangsan",
		Age:   18,
		Score: 99.7,
	}

	//有序初始化并返回Student类型
	// head := Student{
	// 	"zhangsan",
	// 	18,
	// 	99.7,
	// 	nil,
	// }

	//无序初始化并返回指针
	// var head Student
	// head.Name = "zhangsan"
	// head.Age = 18
	// head.Score = 99.7

	// insertTail(&head)
	// insertHead(head)
	delNode(head)
	list(head)

}
