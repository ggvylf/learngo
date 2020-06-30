package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//创建单单层目录
	err := os.Mkdir("./dir1", 0755)
	if err != nil {
		fmt.Println(err)
	}

	//创建多层目录
	err = os.MkdirAll("./dir2/aa/bb/cc", 0755)
	if err != nil {
		fmt.Println(err)
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)

	}
	for _, file := range files {
		fmt.Println(file.Name())

	}
}

