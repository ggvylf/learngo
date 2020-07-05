package main

import (
	"fmt"
	"runtime"
)

func showCaller() {
	//这里的n表示调用层数
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller failed")
	}

	fmt.Println("调用的conter=", pc)
	fmt.Println("调用的文件=", file)
	fmt.Println("调用的行号=", line)

	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println("调用函数的名称=", funcName)

}

func level2() {
	showCaller()
}

func level1() {
	level2()

}

func main() {
	level1()
}
