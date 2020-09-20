package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	for k, v := range os.Args {
		fmt.Printf("第%d个参数=%v\n", k, v)
	}
}
