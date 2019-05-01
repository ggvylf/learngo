package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s1 := "helloworld"

	//是否包含字符串
	fmt.Println(strings.Contains(s1, "o"))

	//是否包含任意一个字符
	fmt.Println(strings.ContainsAny(s1, "O"))

	//前缀
	s2 := "abcdxxxxx1234"
	if strings.HasPrefix(s2, "abcd") {
		fmt.Println("has prefix")
	}

	//后缀s
	if strings.HasSuffix(s2, "1234") {
		fmt.Println("has suffix")
	}
	//返回第一个匹配的索引，没有返回-1
	fmt.Println(strings.Index(s1, "o"))
	//返回最后一个索引
	fmt.Println(strings.LastIndex(s1, "o"))
	//返回字符中的任意一个字符第一次的位置
	fmt.Println(strings.IndexAny(s1, "ld"))

	//字符串拼接
	s3 := []string{"hello", "world"}
	fmt.Println(strings.Join(s3, "%"))

	a := "hehe"
	b := 1234
	str1 := fmt.Sprintf("%s+%d\n", a, b)
	str2 := fmt.Sprintln(a, b)
	str3:=fmt.Sprint(a,b)
	str4 := a + strconv.Itoa(b)
	
	fmt.Println("str1=", str1)
	fmt.Println("str2=", str2)
	fmt.Println("str3=", str3)
	fmt.Println("str4=", str4)

	//切割
	s4 := "hello,world,abc,i,love,you"
	fmt.Println(strings.Split(s4, ","))
	for _, v := range strings.Split(s4, ",") {
		fmt.Println(v)
	}
	fmt.Println(strings.SplitN(s4, ",", 2))
	for _, v := range strings.SplitN(s4, ",", 3) {
		fmt.Println(v)
	}
}
