package main

import "fmt"

func main() {
	str1 := "hello呵呵呵"
	s1 := []rune(str1)
	s2 := make([]rune, 10)
	for _, v := range s1 {
		l := len(string(v))
		switch l {
		case 1:
			fmt.Println("字母=", string(v))
		default:
			fmt.Println("汉字=", string(v))
			s2 = append(s2, v)
		}
	}
	fmt.Println("s2=", string(s2))
}
