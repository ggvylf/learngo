package main

import (
	"errors"
	"fmt"
)

// 创建预制错误类型
var LessThanTwoError error = errors.New("input must not less 2")
var LargeThanThoundError error = errors.New("input largest is 100")

func Fib(n int) ([]int, error) {

	switch {
	case n < 2:
		return nil, LessThanTwoError
	case n > 100:
		return nil, LargeThanThoundError

	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil

}

func main() {

	fib, err := Fib(5)

	// 判断错误类型
	if err == LessThanTwoError {
		fmt.Println("input to less")
	}

	if err == LargeThanThoundError {
		fmt.Println("input to large")
	}
	fmt.Println(fib)
}
