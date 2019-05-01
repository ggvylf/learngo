package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {
	//表达式，支持到秒粒度，年份支持到2099
	expr, err := cronexpr.Parse("*/5 * * * * * *")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(expr)

	//计算下一次调度时间
	now := time.Now()
	nextTime := expr.Next(now)
	fmt.Println(now)
	fmt.Println(nextTime)

	//等待定时器超时后执行
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("run at", nextTime)
	})

	//防止main退出
	time.Sleep(5 * time.Second)

}
