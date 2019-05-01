package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc a7c mfc  cat 8ca azc cba"

	ret := regexp.MustCompile("a.c")

	alls := ret.FindAllStringSubmatch(str, -1)

	fmt.Println(alls)
}
