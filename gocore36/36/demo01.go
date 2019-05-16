package main

import "fmt"

func main() {
	str1 := "GO爱好者"
	str2 := "爱"

	len1 := len(str1)
	fmt.Println(len1)
	len2 := len(str2)
	fmt.Println(len2)

	long(str1)
	long(str2)

}

//方法1 判断单个字符长度
func long(str string) {
	len := len(str)
	if len != 1 && len != 3 {
		fmt.Println(str, "不是单个字符")
	} else {
		fmt.Println(str, "是单个字符")
	}

}
