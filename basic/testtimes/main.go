package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	//当前时间
	fmt.Println(now)

	//当前的年月日时分秒
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//const
	//time.Duration 一纳秒
	//time.Nanosecond 常量，等同于1 time.Duration

	time.Sleep(1 * time.Nanosecond)

	//格式化输出
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//格式化输出，根据参考时间的样式进行输出
	//餐开始键：2016-01-02 15:04:00，这也是go诞生的时间
	fmt.Println(now.Format("2006-01-02 15-04"))
	fmt.Println(now.Format("02-01/2016 15-04-00"))
}
