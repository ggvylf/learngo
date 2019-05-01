package main

import (
	"fmt"
)

type Cats struct {
	Name string
	Age  int
}

//map类型
func TestMapChan() {
	mapChan := make(chan map[string]string, 10)

	m1 := make(map[string]string, 20)
	m1["hehe"] = "haha"
	m1["hello"] = "world"

	mapChan <- m1

	m1out := <-mapChan

	fmt.Println("m1out=", m1out)

}

//struct
func TestSturctChan() {

	cat1 := Cats{"tom", 18}
	cat2 := Cats{"jerry", 13}

	catChan := make(chan Cats, 10)
	catChan <- cat1
	catChan <- cat2

	cat1out := <-catChan
	cat2out := <-catChan

	fmt.Println(cat1out, cat2out)
}

//任意类型，使用interface{}实现
func TestInterfaceChan() {

	interChan := make(chan interface{}, 10)

	in1 := 3
	in2 := "hehe"
	in3 := [3]int{1, 2, 3}
	in4 := Cats{"tom", 13}

	interChan <- in1
	interChan <- in2
	interChan <- in3
	interChan <- in4

	in1out := <-interChan
	in2out := <-interChan
	in3out := <-interChan
	in4out := <-interChan

	//对于struct要使用类型断言才能取到字段
	fmt.Println(in1out, in2out, in3out, in4out.(Cats).Name)

}

func main() {
	TestMapChan()
	TestSturctChan()
	TestInterfaceChan()

}
