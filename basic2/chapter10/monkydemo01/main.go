package main

import "fmt"

type Monkey struct {
	Name string
}

func (m *Monkey) Climbing() {
	fmt.Println(m.Name, "can climbing")
}

//继承
type LittleMonkey struct {
	Monkey
}

//定义一个鸟类的能力的接口
type BirdAble interface {
	Fly()
}

type FishAble interface {
	Swim()
}

//实现接口
func (m *LittleMonkey) Fly() {
	fmt.Println(m.Name, "can fly")
}

func (m *LittleMonkey) Swim() {
	fmt.Println(m.Name, "can swim")
}

func main() {
	m := &LittleMonkey{Monkey{"wukong"}}
	m.Climbing()
	m.Fly()
	m.Swim()
}
