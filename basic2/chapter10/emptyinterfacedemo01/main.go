package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name string
	Id   string
	Age  int
}

//别名，实际上就是Student的slice
type StudentArray []Student

//实现sort.Sort()中的方法
func (p StudentArray) Len() int {
	return len(p)
}
func (p StudentArray) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}
func (p StudentArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	//填充slice
	var stus StudentArray
	for i := 0; i <= 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", rand.Intn(100)),
			Id:   fmt.Sprintf("110%d", rand.Int()),
			Age:  rand.Intn(100),
		}
		stus = append(stus, stu)
	}
	for _, v := range stus {
		fmt.Println(v)
	}
	fmt.Println("_______________")
	//调用Sort() 方法
	sort.Sort(stus)
	for _, v := range stus {
		fmt.Println(v)
	}
}
