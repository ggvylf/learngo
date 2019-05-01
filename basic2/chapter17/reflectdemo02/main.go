package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())

	//Elem()返回接口保存的值的Value封装，或者是指针指向的值的Value封装
	rVal.Elem().SetInt(50)

}

func main() {
	num := 100
	reflect01(&num)
	fmt.Println(num)

}
