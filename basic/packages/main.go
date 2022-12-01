package main

import (
	//匿名导入包，不使用包内部的数据,只要在对应包前加下划线，编译的时候会被编译到可以执行文件中，常见于db驱动

	_ "github.com/go-sql-driver/mysql"

	"fmt"

	//执行的时候关闭 go mod
	//GO111MODULE="off" go run main.go

	// 这个路径是根据GOPATH/src/的内容 go mod是下载到其他目录去了
	"github.com/ggvylf/learngo/basic/packages/add"
)

// init函数可以包含多个
// 不同package里的init函数执行顺序跟导入顺序有关

func init() {
	fmt.Println("main.go init1 func")
}

func init() {
	fmt.Println("main.go init2 func")
}

func main() {
	fmt.Println(add.Add(3, 4))

}
