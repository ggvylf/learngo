package main

import (
	"fmt"
	"io"
	"os"
)

func copy1(srcfile, destfile string) int {
	file1, err := os.Open(srcfile)
	if err != nil {
		fmt.Println(file1)
	}
	defer file1.Close()

	file2, err := os.OpenFile(destfile, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()

	bs := make([]byte, 4096)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)

		if n == 0 || err == io.EOF {
			fmt.Println("copy ok")
			break
		} else if err != nil {
			fmt.Println(err)
		}
		total += n
		file2.Write(bs[:n])
	}
	return total
}

func copy2(srcfile, destfile string) int {
	file1, err := os.Open(srcfile)
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	file2, err := os.OpenFile(destfile, os.O_CREATE|os.O_WRONLY, 0775)
	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()

	n, err := io.Copy(file2, file1)
	return int(n)
}

func main() {
	srcfile := "/etc/hostname"
	destfile := "./a.txt"
	total := copy2(srcfile, destfile)
	fmt.Println(total)

}
