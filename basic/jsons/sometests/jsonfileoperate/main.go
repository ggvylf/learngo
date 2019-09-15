package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Rmb   float64
	Sex   bool
	Hobby []string
}

func jsontofile() {
	p1 := Person{"tom1", 11, 123.45, true, []string{"hehe", "haha"}}
	p2 := Person{"tom2", 11, 123.45, true, []string{"hehe", "haha"}}
	p3 := Person{"tom3", 11, 123.45, true, []string{"hehe", "haha"}}

	persons := make([]Person, 0)
	persons = append(persons, p1, p2, p3)
	file, _ := os.OpenFile("./a.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err := encoder.Encode(persons)
	if err != nil {
		fmt.Println(err)
	}
}

func filetojson() {
	file, _ := os.Open("./a.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	persons := make([]Person, 0)
	err := decoder.Decode(&persons)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range persons {
		fmt.Println(v)
	}

}

func main() {
	jsontofile()
	filetojson()

}
