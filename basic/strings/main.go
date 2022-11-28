package main

import (
	"fmt"
	"strings"
)

func main() {

	//包含字符串
	contains()

	//字符串的索引操作
	index()

	//大小写转换
	upperlower()

	//字符串的修剪
	trim()

	//字符串的分割
	split()

	//字符串数组元素之间加入分隔符
	join()

	//unicode和utf8的差异
	unicodeandutf8()
}

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
	//指定空为分隔符
	strs := strings.Split("hello world i love you", " ")
	fmt.Printf("strs的类型是%T,内容=%s,长度=%d\n", strs, strs, len(strs))

	//指定空格为分隔符，分割为3段，-1表示全部分割
	strs = strings.SplitN("hello world i love you", " ", 3)
	fmt.Printf("strs的类型是%T,内容=%s,长度=%d\n", strs, strs, len(strs))
	for k, v := range strs {
		fmt.Printf("str[%d]=%s\n", k, v)
	}

	//包含分隔符
	strs = strings.SplitAfter("hello,world,i,love,you", ",")
	fmt.Printf("strs的类型是%T,内容=%s,长度=%d\n", strs, strs, len(strs))

	//包含分隔符,分割为制定的段数
	strs = strings.SplitAfterN("hello,world,i,love,you", ",", 3)
	fmt.Printf("strs的类型是%T,内容=%s,长度=%d\n", strs, strs, len(strs))
	for k, v := range strs {
		fmt.Printf("str[%d]=%s\n", k, v)
	}

}

func join() {
	strs := []string{"hello", "world", "i", "love", "you"}
	newstrs := strings.Join(strs, "#")
	fmt.Println(newstrs)

}

func unicodeandutf8() {
	str := "中"
	fmt.Printf(" 中 len is %d\n", len(str))

	//这里是rune切片
	r := []rune(str)
	fmt.Printf(" 中 after []rune  is %x\n", r)
	fmt.Printf(" 中 after []rune len is %d\n", len(r))

	// 遍历中文string 会自动转换成rune类型
	s := "白日依山尽"
	for _, v := range s {

		//fmt.Printf("%c,%d\n", v, v)
		//只使用第一个参数
		fmt.Printf("%[1]c,%[1]d \n", v)
	}

	//这里是byte类型切片
	b := []byte(str)
	fmt.Printf(" 中 after []byes  is %x\n", b)
	fmt.Printf(" 中 after []byes  len is %d\n", len(b))

	fmt.Printf(" 中 unicode is %x\n", r[0])
	fmt.Printf(" 中 utf8 is %x\n", str)

}
