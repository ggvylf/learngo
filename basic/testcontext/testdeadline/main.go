package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//设置deadline
	d := time.Now().Add(4 * time.Second)

	//初始化context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {

	case <-time.After(5 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("deadline ,ext.err=", ctx.Err())
	}

}
