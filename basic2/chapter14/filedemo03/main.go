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

}
