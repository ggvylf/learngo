package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	stu01 := Student{"tom", 18}

	//interface{}转reflect.Value
	rVal := reflect.ValueOf(stu01)

	//reflect.Value转interface{}
	iVal := rVal.Interface()

	//interface{}转Struct，使用类型断言
	stu02 := iVal.(Student)

	fmt.Println(stu02)
	fmt.Println(stu02.Name)
	fmt.Println(stu02.Age)
}
