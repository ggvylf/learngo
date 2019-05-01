package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "someone name")
}

func main() {
	flag.Parse()
	fmt.Printf("hello %s\n", name)
}
