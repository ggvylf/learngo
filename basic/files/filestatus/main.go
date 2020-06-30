package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("/etc/hostname")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info.IsDir())
	fmt.Println(info.ModTime())
	fmt.Println(info.Mode())
	fmt.Println(info.Name())
	fmt.Println(info.Size())
	fmt.Println(info.Sys())

}
