package main

import (
	"fmt"
	"os"
)

//员工信息
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (emp *Emp) ShowMe() {
	fmt.Println("linkid=%d", emp.Id%7)
}

//不带表头的链表
type EmpLink struct {
	Head *Emp
}

//添加员工
func (el *EmpLink) Insert(emp *Emp) {
	//定义辅助指针
	cur := el.Head
	//pre始终再cur的前边
	var pre *Emp = nil

	//el是空链表
	if cur == nil {
		//插入
		el.Head = emp
		return
	}
	for {
		//el不是空链表
		if cur != nil {
			//找到emp插入的位置,如果不满足继续找
			if cur.Id > emp.Id {
				break
			}

			//保证pre永远再cur的前边
			pre = cur
			//移动cur
			cur = cur.Next

		} else {
			//是空链表，直接退出循环
			break
		}
	}

	//退出循环后添加元素，保证pre再cur之前
	if cur == nil {
		pre.Next = emp
		cur.Next = cur
	} else {
		pre.Next = emp
		cur.Next = cur
	}
}

func (el *EmpLink) ShowLink(no int) {
	if el.Head == nil {
		fmt.Println("empty el,no=", no)
		return
	}

	cur := el.Head

	for {
		if cur != nil {
			fmt.Printf("链表%d  id=%d name=%s\n", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
}

func (el *EmpLink) FindById(id int) *Emp {
	cur := el.Head

	for {

		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil

}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (ht *HashTable) ShowAll() {
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].ShowLink(i)
	}
}

//添加员工
func (ht *HashTable) Insert(emp *Emp) {
	linkno := ht.HashFun(emp.Id)
	ht.LinkArr[linkno].Insert(emp)

}

//返回员工存储到哪个链表下标
func (ht *HashTable) HashFun(id int) int {
	return id % 7

}

func (ht *HashTable) FindById(id int) *Emp {
	linkno := ht.HashFun(id)
	return ht.LinkArr[linkno].FindById(id)

}

func main() {
	key := ""
	id := 0
	name := ""
	var hs HashTable

	for {
		fmt.Println("input")
		fmt.Println("show")
		fmt.Println("find")
		fmt.Println("exit")
		fmt.Println("input one")
		fmt.Scanf("%s\n", &key)

		switch key {
		case "input":
			fmt.Println("input id:")
			fmt.Scanf("%d\n", &id)
			fmt.Println("input name:")
			fmt.Scanf("%s\n", &name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hs.Insert(emp)

		case "show":
			hs.ShowAll()
		case "find":
			fmt.Println("input id:")
			fmt.Scanf("%d\n", id)
			emp := hs.FindById(id)
			if emp == nil {
				fmt.Println("not exists")
			}

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("input err")

		}

	}

}
