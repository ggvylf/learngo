package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func BufReadOneLine() {
	//打开文件
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件
	defer file.Close()

	//读取文件
	//创建一个reader
	reader := bufio.NewReader(file)
	//创建一个字符串line存放数据，使用\n作为分隔符。
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("读取到的内容为： %s\n", line)
}

func BufReadAll() {
	//打开文件
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件
	defer file.Close()
	reader := bufio.NewReader(file)
	//读取文件
	for {
		line, err := reader.ReadString('\n')
		//判断是否读取到文件末尾
		if err == io.EOF {
			fmt.Println("已读到文件末尾")
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("读取到的内容为： %s\n", line)
	}

}

func main() {
	BufReadOneLine()
	BufReadAll()

}
