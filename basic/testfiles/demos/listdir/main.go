package main

import (
	"fmt"
	"io/ioutil"
)

func listdir(dirname string) {
	files, _ := ioutil.ReadDir(dirname)

	for _, file := range files {
		filename := dirname + "/" + file.Name()
		fmt.Println(filename)
		if file.IsDir() {
			listdir(filename)
		}
	}
}

func main() {
	dirname := "/tmp"
	listdir(dirname)
}
