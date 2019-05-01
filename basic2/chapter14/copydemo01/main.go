package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(destFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	destFile, err := os.OpenFile(destFileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destFile.Close()
	writer := bufio.NewWriter(destFile)

	return io.Copy(writer, reader)

}

func main() {
	srcFileName := "./a.jpg"
	desFileName := "./b.png"

	_, err := CopyFile(desFileName, srcFileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("拷贝完成")
	}

}
