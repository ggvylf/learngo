package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadFile() {
	//打开文件
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件
	defer file.Close()

	//创建[]byte数组，存放读取的内容。
	tmp := make([]byte, 128)

	//读取文件
	n, err := file.Read(tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("已经读了 %d 字节\n", n)
	fmt.Printf("读取到的内容为： %s\n", string(tmp))
}

func ReadAll() {
	//只读方式打开文件
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件
	defer file.Close()

	//创建[]byte数组，存放读取的内容。
	tmp := make([]byte, 5)
	//读取文件
	for {
		n, err := file.Read(tmp)
		//判断是否读取到文件末尾
		if err == io.EOF {
			fmt.Println("已读到文件末尾")
			fmt.Printf("已经读了 %d 字节\n", n)
			fmt.Printf("读取到的内容为： %s\n", string(tmp[:n]))
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("已经读了 %d 字节\n", n)
		fmt.Printf("读取到的内容为： %s\n", string(tmp[:n]))
	}

}

//使用ioutil来读取文件，注意会读取全部文件，读取大文件不要用这种方法。s
func ReadIoutil(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))

}

func main() {
	ReadFile()
	ReadAll()
	ReadIoutil("./a.txt")

}
