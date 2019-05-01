package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "someone name")
}

func main() {

	flag.Parse()
	hello(name)

}
