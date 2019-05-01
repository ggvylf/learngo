package main

import (
	"errors"
	"fmt"
)

func readFile(name string) (err error) {
	if name == "init.txt" {
		fmt.Println("ok!")
		return nil
	} else {
		return errors.New("filename not matchs")
	}
}

func test() {

	err := readFile("abc.txt")
	fmt.Println("this is test")
	if err != nil {
		panic(err)
	}

}

func main() {
	test()
	fmt.Println("this is main")
}
