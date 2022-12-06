package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", Age: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
