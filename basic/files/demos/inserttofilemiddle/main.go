package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//打开原始文件
	file, err := os.OpenFile("./a.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}

	//打开临时文件，用来存放原始文件中的部分数据
	tmpfile, err := os.OpenFile("./tmp.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	//读取原始文件的第一个字节
	var ret [5]byte
	n, err := file.Read(ret[:])
	if err != nil {
		fmt.Println(err)
	}

	//把读取出的内容写入临时文件
	tmpfile.Write(ret[:n])

	//把要插入的内容插入临时文件
	s := []byte("\nline2")
	tmpfile.Write(s[:])

	//把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := file.Read(x[:])
		if err == io.EOF {
			tmpfile.Write(x[:n])
			fmt.Println("end of file")
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		tmpfile.Write(x[:n])
	}

	file.Close()
	tmpfile.Close()
	os.Rename("./tmp.txt", "./ret.txt")

}
