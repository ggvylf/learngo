package main

import (
	"context"
	"fmt"
	
)

func process(ctx context.Context) {
	//判断是否为int
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		//设定默认值
		ret = 1111111
	}

	fmt.Printf("ret:%d\n", ret)

	s, _ := ctx.Value("session").(string)
	fmt.Printf("session:%s\n", s)

}

func main() {
	//初始化ctx
	ctx := context.WithValue(context.Background(), "trace_id", 13483434)
	//附加到ctx
	ctx = context.WithValue(ctx, "session", "sdlkfjkaslfsalfsafjalskfj")
	process(ctx)

	fmt.Println(ctx)
}
