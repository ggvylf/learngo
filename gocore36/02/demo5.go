package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	cmdLine.StringVar(&name, "name", "everyone", "the greeting object")
}

func main() {
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("hello %s\n", name)
}
