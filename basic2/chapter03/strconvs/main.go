package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Println(str)

	str = fmt.Sprintf("%f", num2)
	fmt.Println(str)

	str = fmt.Sprintf("%t", b)
	fmt.Println(str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Println(str)

	str = strconv.FormatInt(int64(num1), 10)
	fmt.Println(str)

	str = strconv.FormatFloat(num2, 'f', 5, 64)
	fmt.Println(str)

	str = strconv.FormatBool(b)
	fmt.Println(str)

	str2 := strconv.Itoa(num1)
	fmt.Printf("%T %v\n", str2, str2)

	str3, _ := strconv.Atoi(str2)
	fmt.Printf("%T %v\n", str3, str3)

	str4, _ := strconv.ParseInt(str2, 10, 0)
	fmt.Printf("%T %v\n", str4, str4)
}
