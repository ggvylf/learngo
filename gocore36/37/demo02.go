package main

import (
	"strings"
)

func main() {
	// 示例1。
	var builder1 strings.Builder
	//相当于实际使用了
	builder1.Grow(1)

	f1 := func(b strings.Builder) {
		//b.Grow(1) // 这里会引发panic。
	}
	f1(builder1)

	ch1 := make(chan strings.Builder, 1)
	ch1 <- builder1
	builder2 := <-ch1
	//builder2.Grow(1) // 这里会引发panic。
	_ = builder2

	builder3 := builder1
	//builder3.Grow(1) // 这里会引发panic。
	_ = builder3

	// 示例2。
	//传指针
	f2 := func(bp *strings.Builder) {
		(*bp).Grow(1) // 这里虽然不会引发panic，但不是并发安全的。
		//跟上边一样，被使用了就不能复制
		builder4 := *bp
		//builder4.Grow(1) // 这里会引发panic。
		_ = builder4
	}
	f2(&builder1)

	//reset以后归零值
	builder1.Reset()
	//复制以后才是真正使用
	builder5 := builder1
	builder5.Grow(1) // 这里不会引发panic。
}
