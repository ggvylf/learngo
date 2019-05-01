package main

import (
	"fmt"
	"strings"
)

//判读字符串str在url中第一次出现的位置，从0开始，如果不存在返回-1
func firstProcess(url, str string) int {
	result := strings.Index(url, str)
	return result
}

//判断字符串str在url中最后一次出现的位置，从0开始，如果不存在返回-1
func lastProcess(url, str string) int {
	result := strings.LastIndex(url, str)
	return result

}

func main() {
	var str string
	var url string
	fmt.Scanf("%s%s", &url, &str)

	first := firstProcess(url, str)
	last := lastProcess(url, str)

	fmt.Println(first)
	fmt.Println(last)

}
