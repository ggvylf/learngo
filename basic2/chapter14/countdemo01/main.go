package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {

	fileName := "/etc/passwd"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var cc CharCount

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				cc.ChCount++
			case v == ' ' || v == '\t':
				cc.SpaceCount++
			case v >= 0 && v <= 9:
				cc.NumCount++
			default:
				cc.OtherCount++
			}

		}
	}
	fmt.Println(cc)

}
