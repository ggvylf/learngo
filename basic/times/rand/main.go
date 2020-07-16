package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		r := rand.Intn(10)
		fmt.Println(r)
	}

	//返回随机时间
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
}
