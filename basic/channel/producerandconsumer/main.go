package main

import (
	"fmt"
	"strconv"
	"time"
)

type Product struct {
	Name string
}

func Producer(chanStorage chan<- Product, chanQuit chan<- interface{}) {

	for {
		if Flag {
			fmt.Println("Producer get Flag")
			chanQuit <- "go die"
		}
		product := Product{
			Name: "item" + strconv.Itoa(time.Now().Nanosecond()),
		}
		chanStorage <- product
		fmt.Println("sent to Storage item=", product)
		time.Sleep(1 * time.Second)
	}

}

func Consumer(chanShop <-chan Product) {
	for {
		product := <-chanShop
		fmt.Println("rec from shop item=", product)

	}

}

func Logistics(chanStorage <-chan Product, chanShop chan Product) {

	for {
		product := <-chanStorage
		chanShop <- product
	}

}

var Flag bool

func Watcher() {

	for {
		sec := time.Now().Second()
		fmt.Println("now second=", sec)
		if sec == 0 || sec == 30 {
			Flag = true
			fmt.Println("sec is over")
			return
		}
		time.Sleep(1 * time.Second)
	}

}

func main() {

	chanStorage := make(chan Product, 10)
	chanShop := make(chan Product, 5)
	chanQuit := make(chan interface{})

	go Producer(chanStorage, chanQuit)
	go Logistics(chanStorage, chanShop)
	go Consumer(chanShop)
	go Watcher()

	<-chanQuit

	fmt.Println("main func end")

}
