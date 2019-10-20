package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func MyEncode(filename string) error {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer file.Close()
	encoder := json.NewEncoder(file)
	p := MyDecode(Srcfile)
	err := encoder.Encode(p)
	if err != nil {
		fmt.Println("编码失败", err)
		return err
	} else {
		fmt.Println("编码成功")
		return nil
	}

}

func MyDecode(filename string) *Person {

	file, _ := os.Open(filename)

	defer file.Close()
	decoder := json.NewDecoder(file)

	p := new(Person)
	err := decoder.Decode(p)
	if err != nil {

		fmt.Println("解码失败", err)
	} else {
		fmt.Println("解码成功")
	}

	return p

}

var Srcfile = "./info.json"
var Destfile = "./a.txt"

func main() {

	MyEncode(Destfile)
}
