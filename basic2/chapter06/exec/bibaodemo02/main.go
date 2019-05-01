package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if !strings.HasSuffix(filename, suffix) {
			return filename + suffix

		}
		return filename
	}

}

func main() {
	a := makeSuffix(".jpg")
	fmt.Println(a("abc"))
	b:=makeSuffix(".aaa")
	fmt.Println(b("c.jpg"))
}
