package main

import (
	"fmt"
	"os"
)

func PathExits(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func main() {
	path := "./b.txt"

	result, err := PathExits(path)
	if result == true && err == nil {
		fmt.Println("文件存在")
	}
	if result == false || err != nil {
		fmt.Println("文件不存在，错误为", err)
	}

}
