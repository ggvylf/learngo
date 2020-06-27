//可以看到append()的操作，cap默认增加2倍
package main

import (
	"fmt"
)

func main() {
	//声明一个arr
	var a [5]int = [...]int{1, 2, 3, 4, 5}

	//对arr进行切片
	slice := a[:]
	fmt.Println(slice)

	fmt.Printf("the slice memory address is %p\nslice is %v\nlen is %d\ncap is %d\n", &slice, slice, len(slice), cap(slice))

	slice = append(slice, 10)
	fmt.Printf("after append,the slice memory address is %p\nslice is %v\nlen is %d\ncap is %d\n", &slice, slice, len(slice), cap(slice))

	slice = append(slice, 10)
	fmt.Printf("after append,the slice memory address is %p\nslice is %v\nlen is %d\ncap is %d\n", &slice, slice, len(slice), cap(slice))

	slice = append(slice, 10)
	fmt.Printf("after append,the slice memory address is %p\nslice is %v\nlen is %d\ncap is %d\n", &slice, slice, len(slice), cap(slice))

	//声明slice
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	//slice的追加，会自动扩容。
	s3 := append(s1, s2...)
	fmt.Println(s3)

	//slice的拷贝，按照s4的长度为准，如果s3比s4大，则会copy一部分s3，不会扩容
	s4 := make([]int, 10)
	copy(s4, s3)
	fmt.Println(s4)

	//string底层是一个[]byte数组
	str := "hello world"
	str1 := str[:5]
	str2 := str[6:]
	fmt.Println(str1)
	fmt.Println(str2)

	//改变string中的字符
	str3 := []byte(str)
	str3[2] = 'k'
	str = string(str3)
	fmt.Println(str)

}
