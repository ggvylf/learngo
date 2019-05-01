package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "/etc/passwd"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read end")
			break
		}
		fmt.Printf("%v", str)

	}
}
