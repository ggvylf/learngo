//判断一个字符串是不是回文
//例如 123321 abcba 人过大佛寺寺佛大过人
package main

import "fmt"

func huiwen(n string) {
	//这里转换为[]rune来支持中文
	str := []rune(n)
	flag := false
	for i := 0; i < len(str); i++ {
		if str[i] == str[len(str)-i-1] {
			continue
		} else {
			flag = true
			break
		}

	}
	if flag {
		fmt.Println("not ok")
	} else {
		fmt.Println("ok")
	}

}

func main() {
	var str string
	//fmt.Scanf("%s", &str)
	str = "人过大佛寺寺佛大过人"
	huiwen(str)

}
