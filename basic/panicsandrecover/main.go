package main

//painc 用于不可恢复的错误
//painc 在退出前会执行defer指定的内容
// 可以在defer函数中使用recover()函数来避免程序退出

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func area(r float64) float64 {
	if r <= 0 {
		panic(errors.New("半径不能为负数"))

	}
	return math.Pi * r * r

}

// os.Exit退出的时候不会调用defer
// os.Exit退出的时候不输出当前调用栈信息

func areawithosexit(r float64) float64 {
	if r <= 0 {
		fmt.Println("半径不能为负数")
		os.Exit(0)
	}

	return math.Pi * r * r

}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic reason is ", err)
			fmt.Println("recover run succ")
		}

	}()

	r, _ := strconv.Atoi(os.Args[1])

	res := area(float64(r))
	fmt.Println("面积=", res)

	//oe.Exit没有调用栈信息
	// res2 := areawithosexit(float64(r))
	// fmt.Println("面积2=", res2)
}
