package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	// flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	flag.CommandLine.StringVar(&name, "name", "everyone", "the greeting object")
}

func main() {
	flag.Parse()
	fmt.Printf("hello %s\n", name)
}
