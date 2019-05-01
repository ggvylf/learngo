package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	bs := []byte{'a', 'b', 'c', 'd'}

	n, err := file.Write(bs)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	n, err = file.WriteString("hello world")
	
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(n)


}
