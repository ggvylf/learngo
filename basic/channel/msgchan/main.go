package main

import "fmt"

func count(name string, i int, msgChan chan int) {

	for j := 0; j < i; j++ {

		fmt.Printf("%s=%d\n", name, j)
	}
	fmt.Println(name, "end")
	msgChan <- 0
}

func main() {

	msgChan := make(chan int, 3)

	go count("tom", 3, msgChan)
	go count("jerry", 2, msgChan)
	count("main", 4, msgChan)

	for i := 0; i < 3; i++ {
		<-msgChan
	}

	fmt.Println("main func end")
}
