package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	mych := make(chan bool)
	oncefunc := func() {
		fmt.Println("only run once")
	}

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(oncefunc)
			fmt.Println("in goroutine")
			mych <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-mych
	}
}
