/*
go run --gcflags '-m -m -l' dynamicEscape.go
# command-line-arguments
./dynamicEscape.go:6:14: "hello world" escapes to heap:
./dynamicEscape.go:6:14:   flow: {storage for ... argument} = &{storage for "hello world"}:
./dynamicEscape.go:6:14:     from "hello world" (spill) at ./dynamicEscape.go:6:14
./dynamicEscape.go:6:14:     from ... argument (slice-literal-element) at ./dynamicEscape.go:6:13
./dynamicEscape.go:6:14:   flow: {heap} = {storage for ... argument}:
./dynamicEscape.go:6:14:     from ... argument (spill) at ./dynamicEscape.go:6:13
./dynamicEscape.go:6:14:     from fmt.Println(... argument...) (call parameter) at ./dynamicEscape.go:6:13
./dynamicEscape.go:6:13: ... argument does not escape
./dynamicEscape.go:6:14: "hello world" escapes to heap
hello world
*/

package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
