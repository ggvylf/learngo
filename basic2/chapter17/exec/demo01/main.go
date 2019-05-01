package main

import (
	"fmt"
	"reflect"
)

func refTest(i interface{}) {
	rType := reflect.TypeOf(i)
	fmt.Println("Kind=", rType.Kind())

	rVal := reflect.ValueOf(i)
	fmt.Println("Type=", rVal.Type())
	fmt.Println("Value=", rVal.Float())

	iV := rVal.Interface()

	fmt.Println("iV=", iV.(float64))

}

func refEdit(i interface{}) {
	rVal := reflect.ValueOf(i)
	rVal.Elem().SetFloat(4.2)
	fmt.Println("after edit=", rVal.Elem().Float())

}

func main() {
	num := 1.2
	refTest(num)
	refEdit(&num)

	fmt.Println("main num=", num)

}
