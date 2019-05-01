package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.OpenFile("./a.txt", os.O_CREATE|os.O_RDWR, 0755)
	str := "abcdefg"
	file.WriteString(str)
	defer file.Close()

	bs := []byte{0}

	//从文件开头光标偏移1个
	file.Seek(1, 0)
	file.Read(bs)
	fmt.Println(string(bs))

	//从当前位置光标偏移1个
	file.Seek(2, 1)
	file.Read(bs)
	fmt.Println(string(bs))

	//从文件尾部光标向左偏移2个
	file.Seek(-2, 2)
	file.Read(bs)
	fmt.Println(string(bs))
}
