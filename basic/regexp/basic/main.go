package main

import (
	"fmt"
	"regexp"
	"strings"
)

func myregexp() {
	rerule := `[\s\S]+@[\s\S]+[ ]?`
	str := "1@1.com 2@2.net 3@3.ml"
	re := regexp.MustCompile(rerule)
	result := re.FindAllString(str, -1)
	for _, rst := range strings.Split(result[0], " ") {
		fmt.Println(rst)
	}
}

func main() {
	myregexp()
}
