package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func read1() {
	file, err := os.OpenFile("./a.txt", os.O_RDONLY, 0644)
	defer func() {
		file.Close()
		fmt.Println("文件正常关闭")
	}()

	if err != nil {
		fmt.Println("读取文件失败,err=", err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == nil {
			fmt.Println(str)
		} else {
			if err == io.EOF {
				fmt.Println("到达文件末尾")
				break
			}
			fmt.Println("读取文件失败,err=", err)
			return
		}
	}
}

func read2() {
	bytes, err := ioutil.ReadFile("./a.txt")
	if err == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("文件读取错误，err=", err)
		return
	}
	fmt.Println("文件读取完成")
}

func write1() {
	file, _ := os.OpenFile("./b.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("aaaa\n")
	writer.WriteString("bbbb\n")
	writer.Flush()
}

func copy1() {
	src, _ := os.OpenFile("./a.txt", os.O_RDONLY, 0644)
	dest, _ := os.OpenFile("./c.txt", os.O_CREATE|os.O_WRONLY, 0644)
	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dest)
	buf := make([]byte, 1024)

	defer func() {
		src.Close()
		dest.Close()
		fmt.Println("all closed!")
	}()

	for {
		_, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("已读取到文件末尾")
				break
			} else {
				fmt.Println("读取错误，err=", err)
				return
			}

		} else {
			_, err := writer.Write(buf)
			if err != nil {
				fmt.Println("写入错误，err=", err)
				return
			}
		}

	}
	fmt.Println("copy ok!")
}

func main() {
	//read2()
	//write1()
	copy1()

}
