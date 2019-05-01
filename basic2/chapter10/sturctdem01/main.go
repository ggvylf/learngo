package main

import (
	"fmt"
)

type Cats struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

func (cat Cats) Call() {
	fmt.Println("miaomiao")
}

func (cat Cats) Rename() {
	cat.Name = "tom"
	fmt.Println("Rename cat=", cat)
}

func main() {
	cat1 := Cats{
		"jerry",
		2,
		"black",
	}

	cat1.Call()
	cat1.Rename()
	fmt.Println("cat=", cat1)

}
