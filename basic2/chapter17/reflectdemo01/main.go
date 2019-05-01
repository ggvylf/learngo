package main

import (
	"fmt"
	"reflect"
)

type Studnet struct {
	Name string
	Age  int
}

func reflectTestStuct(i interface{}) {

	rType := reflect.TypeOf(i)
	fmt.Println(rType)

	rVal := reflect.ValueOf(i)
	fmt.Println(rVal)

	iV := rVal.Interface()

	//类型断言
	stu, ok := iV.(Studnet)
	if ok {
		fmt.Println("stu.Name=", stu.Name)
		fmt.Println("stu.Age=", stu.Age)
	}

}

func main() {
	stu01 := Studnet{"tom", 18}
	reflectTestStuct(stu01)
}
