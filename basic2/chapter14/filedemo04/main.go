package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./a.txt"
	
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	str := "hello world\n"

	//带缓冲区的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//从缓存存写入磁盘
	writer.Flush()

}
