package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	//从字符串中解析出对应的数据
	//base表示转换的进制，0表示不转换，bitSize表示转换后是int8 int16 int32，0表示转换为int类型
	i, err := strconv.ParseInt(str, 10, 8)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("i=%v,type=%T\n", i, i)

	//字符串转换为int类型
	//等同于strconv.Parse(str,0.0)
	i2, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("i2=%v,type=%T\n", i2, i2)

	//数字转字符串
	s1 := strconv.Itoa(i2)

	fmt.Printf("s1=%v,type=%T\n", s1, s1)

	//从字符串解析bool值
	str2 := "true"
	b1, err := strconv.ParseBool(str2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("b1=%v,type=%T\n", b1, b1)

	//从字符串解析出float
	str3 := "1.234"
	f1, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("f1=%v,type=%T\n", f1, f1)

}
