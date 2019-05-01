//冒泡排序
package main

import "fmt"

func maoPao(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
	fmt.Println(a)

}

func main() {
	a := []int{1, 22, 7, 9, 55, 3, 4}
	fmt.Println(a)
	maoPao(a)
}
