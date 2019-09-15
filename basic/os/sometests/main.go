package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func copyfile() {
	src, _ := os.OpenFile("./a.txt", os.O_RDONLY, 0644)
	dest, _ := os.OpenFile("./c.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	defer func() {
		src.Close()
		dest.Close()
		fmt.Println("文件已关闭")
	}()
	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dest)

	buf := make([]byte, 1024)

	for {
		n, err := reader.Read(buf)
		fmt.Println("已读取=", n)
		if err == nil {
			m, err := writer.Write(buf)
			fmt.Println("已写入=", m)
			if err != nil {
				fmt.Println("写入错误，err=", err)
				return
			} else {
				fmt.Println("写入成功")
			}
		} else {

			if err == io.EOF {
				fmt.Println("已读取到文件末尾")
				break
			} else {
				fmt.Println("读取文件错误，err=", err)
				return
			}
		}

	}
	writer.Flush()
	fmt.Println("拷贝完毕")

}

func copyioutil() {
	bytes, _ := ioutil.ReadFile("./a.txt")
	err := ioutil.WriteFile("./c.txt", bytes, 0644)
	if err == nil {
		fmt.Println("文件复制成功")
	} else {
		fmt.Println("文件复制失败，err=", err)
	}
}

func iocopy() {
	src, _ := os.OpenFile("./a.txt", os.O_RDONLY, 0644)
	dest, _ := os.OpenFile("./c.txt", os.O_WRONLY|os.O_CREATE, 0644)

	defer func() {
		src.Close()
		dest.Close()
		fmt.Println("文件已关闭")
	}()
	written, err := io.Copy(dest, src)
	if err == nil {
		fmt.Println("文件复制成功,已赋值byte数=", written)
	} else {
		fmt.Println("文件复制失败，err=", err)
	}

}

func main() {
	copyfile()
	//copyioutil()
	//iocopy()
}
