package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	srcfile := "/etc/passwd"
	destfile := "/tmp/hehe"
	tempfile := "/tmp/hehe.tmp"
	file1, _ := os.Open(srcfile)
	file2, _ := os.OpenFile(destfile, os.O_CREATE|os.O_WRONLY, 0775)
	file3, _ := os.OpenFile(tempfile, os.O_CREATE|os.O_WRONLY, 0775)

	defer file1.Close()
	defer file2.Close()

	bs := make([]byte, 4096)
	n1, err := file3.Read(bs)
	countStr := string(bs[:n1])
	fmt.Println(countStr)

	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println(count)

	file1.Seek(count, 0)
	file2.Seek(count, 0)
	data := make([]byte, 4096)
	n2 := -1
	n3 := -1
	total := int(count)
	for {
		n2, err = file1.Read(data)
		if err == io.EOF {
			fmt.Println("copy ok!")
			break
			file3.Close()
			os.Remove(tempfile)
		}
		n3, _ = file2.Write(data[:n2])
		total += n3
		file3.Seek(0, 0)
		file3.WriteString(strconv.Itoa(total))
	}

}
