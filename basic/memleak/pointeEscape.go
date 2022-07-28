/*
go run --gcflags '-m -m -l' pointeEscape.go
# command-line-arguments
./pointeEscape.go:10:2: sum escapes to heap:
./pointeEscape.go:10:2:   flow: ~r0 = &sum:
./pointeEscape.go:10:2:     from &sum (address-of) at ./pointeEscape.go:11:9
./pointeEscape.go:10:2:     from return &sum (return) at ./pointeEscape.go:11:2
./pointeEscape.go:10:2: moved to heap: sum
*/

package mains

func main() {
	pointerEscape(1, 2)
}

func pointerEscape(a, b int) *int {
	sum := a + b
	return &sum
}
