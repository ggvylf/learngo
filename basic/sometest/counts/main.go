//读取字符串，并统计数字、字母、空格、符号等
package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(str string) (wordCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}

	}
	return

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Printf("word is %d\nspace is %d\nnumber is %d\nother is %d\n", wc, sc, nc, oc)

}
