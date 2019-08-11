package main

import (
	//匿名导入包，不使用包内部的数据,只要在对应包前加下划线，编译的时候会被编译到可以执行文件中，常见于db驱动
	_ "github.com/go-sql-driver/mysql"

	"fmt"

	"github.com/ggvylf/learngo/basic/packages/add"
)

func main() {
	fmt.Println(add.Add(3, 4))

}
