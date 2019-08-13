package main

import (
	"fmt"
	"time"
)

func main() {

	//获取当前时间
	now := time.Now()
	fmt.Printf("当前时间为：%v\n", now)

	//当前的年月日时分秒
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳，返回从1970.01.01到现在的秒数
	fmt.Printf("时间戳：%v 秒\n", now.Unix())
	fmt.Printf("时间戳：%v 纳秒\n", now.UnixNano())

	//把时间戳抓暖胃具体的时间格式
	t := time.Unix(1565689034, 0)
	fmt.Println("时间戳转时间格式：", t)

	//常见的const
	//time.Duration 一纳秒
	//time.Nanosecond 纳秒，等同于1 time.Duration
	//time.Microsecond 微秒 1000*time.Nanosecond
	//time.Second 秒
	//time.Minute 分钟
	//time.Hour 小时

	//时间间隔类型
	//time.Duration
	time.Sleep(1 * time.Nanosecond)
	//不能用int类型直接写在Sleep中。需要做类型转换
	n := 5
	time.Sleep(time.Duration(n) * time.Nanosecond)

	//格式化输出
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//格式化输出，根据参考时间的样式进行输出
	//格式化初始时间：2016-01-02 15:04:00，这也是go诞生的时间
	fmt.Println(now.Format("2006-01-02 15-04"))
	fmt.Println(now.Format("02-01/2016 15-04-00"))
}
