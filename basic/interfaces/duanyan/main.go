//类型断言
package main

import "fmt"

type Coder struct{}

func (c *Coder) Work() {
	fmt.Println("is working")
}

func (c *Coder) Rest() {
	fmt.Println("is resting")
}

type Worker interface {
	Work()
	Rest()
}

func Test(a interface{}) {
	//断言，a是个int类型，如果不是int类型就走if流程
	b, ok := a.(int)
	if !ok {
		fmt.Println("not a int")
		return
	}
	b += 3
	fmt.Printf("%d\n", b)
}

func testType(items ...interface{}) {
	for _, v := range items {
		//使用switch case来进行类型的判断
		switch v.(type) {
		case int, int32, int64:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("unknow")
		}
	}
}

func ptrduanyan() {

	//接口的实例只能使用指针
	var worker Worker
	c := new(Coder)
	worker = c

	//注意这里，对借口类型最断言，只能断言为实例或者是指针中的一种，否则编译器会报错。
	switch worker.(type) {
	//case Coder :
	//	fmt.Println("type is *Coder")
	case *Coder:
		fmt.Println("type is Coder")
	}

}

func main() {
	var a int

	Test(a)

	testType("123", 123, 46)

	ptrduanyan()
}
