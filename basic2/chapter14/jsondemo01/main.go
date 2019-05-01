package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"username"`
	Age  int    `json:"myage"`
}

func TestStruct() {
	stu1 := Student{"tom", 18}

	result, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

}

func TestMap() {
	map1 := make(map[int]string)
	map1[1] = "hehe"
	result, err := json.Marshal(map1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

}

func TestSlice() {
	var s []map[int]int

	m1 := make(map[int]int)
	m1[1] = 1
	m1[2] = 2

	s = append(s, m1)

	m2 := make(map[int]int)
	m2[3] = 3
	m2[4] = 4
	s = append(s, m2)

	result, _ := json.Marshal(s)
	fmt.Println(string(result))

}

func TestFloat64() {
	a := 3.14
	result, _ := json.Marshal(a)
	fmt.Printf("%v,%T", string(result), result)

}

func main() {
	TestStruct()
	TestMap()
	TestSlice()
	TestFloat64()

}
