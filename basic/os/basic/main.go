package main

import (
	"fmt"
	"os"
)

func getdirs() {
	//获取程序运行的工作路径
	dir, _ := os.Getwd()
	fmt.Println("当前工作目录=", dir)
}

func getenv() {
	//遍历环境变量
	//for k,v:=range os.Environ() {
	//	fmt.Printf("第%d个环境变量=%s\n",k+1,v)
	//}

	//获取指定的环境变量
	fmt.Println("环境变量GOPATH的值=", os.Getenv("GOPATH"))

}

func getfilestat() {
	filename := "./main.go"
	stat, _ := os.Stat(filename)
	fmt.Println("文件名=", stat.Name())
	fmt.Println("是否为目录=", stat.IsDir())
	fmt.Println("文件权限=", stat.Mode())
	fmt.Println("文件的最后修改时间=", stat.ModTime())
	fmt.Println("文件大小=", stat.Size())
	fmt.Println("文件的系统参数=", stat.Sys())

}

func main() {
	getdirs()
	getenv()
	getfilestat()
}
