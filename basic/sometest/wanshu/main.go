//输入一个数，判断1到这个数之间有多少是完数

package main

import (
	"fmt"
)

func perfect(n int) bool {
	sum := 0

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return n == sum

}

func process(n int) {
	for i := 1; i < n+1; i++ {
		if perfect(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	process(n)
}
