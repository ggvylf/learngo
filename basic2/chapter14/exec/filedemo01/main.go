package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "./a.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", string(content))

	str := []byte("hehe")
	err = ioutil.WriteFile(file, str, 0644)
	if err != nil {
		fmt.Printf("写入失败:%v\n", err)
	}

}
