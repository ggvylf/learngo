package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := " spaces "
	head := " sp"
	tail := "ces "

	// 去掉str前后的空白字符
	fmt.Println(strings.TrimSpace(str1))

	//去掉str首位的字符串
	fmt.Println(strings.Trim(str1, "s"))

	//去掉str左边的head字符串
	fmt.Println(strings.TrimLeft(str1, head))

	//去掉str右边的tail字符串
	fmt.Println(strings.TrimRight(str1, tail))

	str2 := "hello world hello go"

	//返回空格分割的slice
	fmt.Println(strings.Fields(str2))

	//分割字符串，返货slice
	fmt.Println(strings.Split(str2, " "))

	str3 := []string{"hello", "world", "hello", "go"}

	//用字符串链接
	fmt.Println(strings.Join(str3, ","))

	int := 123

	//int转换为string
	str4 := strconv.Itoa(int)

	fmt.Printf("%s\n", str4)

	//string转换为int
	num, _ := strconv.Atoi(str4)
	fmt.Printf("%d\n", num)
}
