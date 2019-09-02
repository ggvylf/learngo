package main

import (
	"fmt"
	"strings"
)

func contains() {
	str1 := "hello world"
	str2 := "le"
	str3 := "heheaaa"
	str4 := "爱我中华helloworld"
	str5 := rune('我')

	//输出某个字符的unicode
	fmt.Printf("我的unicode=%U\n", str5)

	//输出某个unicode对应的字符
	fmt.Printf("0x6211对应的字符=%c\n", 0x6211)

	//是否包含子串
	fmt.Println("hello world是否包含le", strings.Contains(str1, str2))

	//是否包含任意一个char
	fmt.Println("heheaaa是否包含le中的任意字符", strings.ContainsAny(str3, str2))

	//是否包含任意字符
	fmt.Println("爱我中华helloworld是否包含我", strings.ContainsRune(str4, str5))
	fmt.Println("爱我中华helloworld是否包含我", strings.ContainsRune(str4, 0x6211))

}

func index() {
	str1 := "hello world"
	str2 := "l"
	str3 := "heheaaa"
	str4 := "爱我中华helloworld"
	str5 := rune('我')

	//第一个匹配的index
	fmt.Println("hello worldz中第一个l的index=", strings.Index(str1, str2))
	//最后一个匹配的index
	fmt.Println("hello worldz中最后一个l的index=", strings.LastIndex(str1, str2))

	//是否包含任意一个char
	fmt.Println("heheaaa包含e中的第一个index=", strings.IndexAny(str3, "e"))
	fmt.Println("heheaaa包含e中的最后一个index=", strings.LastIndexAny(str3, "e"))

	//不包含返回-1
	fmt.Println("heheaaa包含l中的第一个index=", strings.IndexAny(str3, "l"))

	//是否包含任意字符
	fmt.Println("爱我中华helloworld中第一个我的index=", strings.IndexRune(str4, str5))
	fmt.Println("爱我中华helloworld中最后一个我的index=", strings.LastIndexAny(str4, str2))
}

func upperlower() {
	fmt.Println("abc转大写=", strings.ToUpper("abc"))
	fmt.Println("ABC转小写=", strings.ToLower("ABC"))
	fmt.Println("abc转为大写标题=", strings.ToTitle("abc"))
	fmt.Println("abc的首字母大写=", strings.Title("abc"))
}

func filter(char rune) bool {
	if char == 'c' || char == 'b' {
		return true
	} else {
		return false
	}
}
func trim() {
	fmt.Println("去掉头尾空格=", strings.TrimSpace(" abc "))
	fmt.Println("去掉abc中前缀为a的字符=", strings.TrimPrefix("abc", "a"))
	fmt.Println("去掉abc中后缀为c的字符=", strings.TrimSuffix("abc", "c"))
	fmt.Println("去掉头尾包含a或c字符的内容=", strings.Trim("abc abc abc", "ab"))

	//filter函数返回true表示要进行trim操作，返回false表示不进行
	fmt.Println("使用trimfunc去掉头尾符合条件的字符=", strings.TrimFunc("bvabc abc abc", filter))
}

func split() {

}

func main() {
	contains()
	index()
	upperlower()
	trim()
	split()
}
