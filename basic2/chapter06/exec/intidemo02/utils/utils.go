package utils

import "fmt"

var Name string
var Age int

func init() {
	Name = "tom"
	Age = 18
	fmt.Println("utils()")
	fmt.Println(Name, Age)
}
