package main

import (
	"fmt"
	"time"
)

func AddTime() {
	//时间+时间间隔
	t := time.Now()
	t2 := t.Add(time.Hour)
	fmt.Println("当前时间t=", t)
	fmt.Println("增加一小时后，t2=", t2)
}

func SubTime() {
	t := time.Now()
	t2 := t.Add(time.Hour)
	//时间之差
	t3 := t.Sub(t2)
	fmt.Println("t2-t时间之差=", t3)
}

func EqualTime() {
	//判断时间是否相等。会考虑时区的影响，不同时区的标准时间爱你也可以比较。如果直接判断==，则还会比较时区和地点
	//定义时区偏移量
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	//指定使用北京时间
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	//定义2个time，时区不同
	d1 := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	d2 := time.Date(2000, 2, 1, 20, 30, 0, 0, beijing)

	//直接比较会比较时区，地区等
	datesEqualUsingEqualOperator := d1 == d2
	//Equal只判断时间
	datesEqualUsingFunction := d1.Equal(d2)

	fmt.Printf("直接使用==比较的结果 = %v\n", datesEqualUsingEqualOperator)
	fmt.Printf("使用Equal比较的结果 = %v\n", datesEqualUsingFunction)

}

func BeforeAndAfter() {
	u := time.Now()
	t2 := u.Add(time.Hour)
	fmt.Println("u=", u)
	fmt.Println("t2=", t2)
	//比较时间点，如果t2表示的时间点在u之前，返回真
	if t2.Before(u) {
		fmt.Println("t2前u后")
	} else {
		fmt.Println("u前t2后")
	}

	//比较时间点，如果t2表示的时间点在u之后，返回真
	if t2.After(u) {
		fmt.Println("u前t2后")
	} else {
		fmt.Println("t2前u后")
	}
}

func TimeTick() {
	//定时器，返回值是一个chan
	//循环输出chan中的内容
	i := 0
	//timer:=time.NewTicker(time.Second)
	//for {
	//	time:=<-timer.C
	//	fmt.Println(time)
	//	i++
	//	if i==10 {
	//		return
	//	}
	//}
	for time := range time.Tick(time.Second) {
		fmt.Println(time)
		i++
		if i == 10 {
			return
		}
	}
}

func main() {
	AddTime()
	SubTime()
	EqualTime()
	BeforeAndAfter()
	TimeTick()
}
