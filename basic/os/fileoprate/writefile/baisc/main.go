package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WriteFile() {
	//打开文件，制定文件的名称，操作的权限和文件的权限
	file, err := os.OpenFile("./testfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	//最后关闭文件
	defer file.Close()

	str := "abcdasdfalksdjg;laskdffasd\n"

	//写入[]byte
	file.Write([]byte(str))

	//写入字符串
	file.WriteString(str)
}

//使用ioutil写入文件
func IOutilWriteFile() {
	str := "sdfffffffffffffasdfasdf"
	err := ioutil.WriteFile("./ioutilsfile.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	WriteFile()
	IOutilWriteFile()

}
