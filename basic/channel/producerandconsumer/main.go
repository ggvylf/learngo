package main

import (
	"fmt"
	"strconv"
	"time"
)

type Product struct {
	Name string
}

func Producer(chanShop chan<- Product) {

	for {
		product := Product{
			Name: "item" + strconv.Itoa(time.Now().Nanosecond()),
		}
		chanShop <- product
		fmt.Println("sent to shop item=", product)
		time.Sleep(1 * time.Second)
	}

}

func Consumer(chanShop <-chan Product, chanCount chan<- int) {

	for {
		product := <-chanShop
		fmt.Println("rec from shop item=", product)
		chanCount <- 1
	}

}

func main() {
	chanShop := make(chan Product, 5)
	chanCount := make(chan int, 5)

	go Producer(chanShop)
	go Consumer(chanShop, chanCount)

	for i := 0; i < 10; i++ {
		<-chanCount
	}

	fmt.Println("main func end")

}
