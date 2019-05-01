package main

import (
	"fmt"
	"math/rand"
<<<<<<< HEAD
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

=======
)

func main() {
>>>>>>> 63bdc97f0367fd8589bda9c86f13de8fca5c9961
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}
