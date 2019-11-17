package main

import (
	"fmt"
	"time"
)

func myticker() {
	flag := false
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		time.Sleep(10 * time.Second)
		ticker.Stop()
		flag = true
	}()
	for {

		if flag {
			break

		}
		fmt.Println(<-ticker.C)
	}
	fmt.Println("myticker end")
}

func main() {
	myticker()
	fmt.Println("main func end")

}
