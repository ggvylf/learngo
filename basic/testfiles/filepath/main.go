package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	file1 := "/etc/hostsname"
	file2:="./a.txt"

	//是否绝对路径
	fmt.Println(filepath.IsAbs(file1))
	fmt.Println(filepath.IsAbs(file2))

	//显示绝对路径
	fmt.Println(filepath.Abs(file1))
	fmt.Println(filepath.Abs(file2))

}
