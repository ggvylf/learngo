package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		//指定文件读取的结束符号，\n表示换行
		str, err := reader.ReadString('\n')

		//io.EOF表示文件的结尾
		if err == io.EOF {
			break
		}
		fmt.Printf("%v", str)
	}

	fmt.Println("文件读取结束")

}
