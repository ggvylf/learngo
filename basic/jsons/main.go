package main

import (
	"encoding/json"
	"fmt"
)

// struct的tag
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {

	stu := Student{
		Name:  "zhangsan",
		Age:   18,
		Score: 18,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}
