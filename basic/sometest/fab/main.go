//斐波那契数列

package main

import "fmt"

//递归的方法
func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

//使用数组
func fab2(n int) {
	var a []int
	a = make([]int, n)

	a[0] = 1
	a[1] = 1
	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	for _, v := range a {
		fmt.Println(v)
	}

}

func main() {
	// var n int
	// fmt.Scanf("%d", &n)
	for i := 1; i < 10; i++ {
		fmt.Println(fab(i))
	}

	fab2(10)

}
