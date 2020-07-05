package main

import (
	"bufio"
	"fmt"
	"os"
)

func BufWriteFile() {
	file, err := os.OpenFile("./bufile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	str := "sdffffffffffasdfasdff\n"

	//定义一个带缓冲区的write
	writer := bufio.NewWriter(file)

	//写入缓冲区
	writer.WriteString(str)

	//把缓冲区中的内如刷新到文件中
	writer.Flush()

}

func main() {
	BufWriteFile()

}
