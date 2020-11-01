package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 建议主动调用下cancel()来释放上下文和父上下文的资源，不要等ctx自己过期
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("hehe")
	case <-ctx.Done():
		fmt.Println("ctx done")
		fmt.Println(ctx.Err())
	}

}
