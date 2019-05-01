package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	//建立连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.56.11:2379"},
		DialTimeout: 5 * time.Second,
	})

	//错误处理
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("context deadline")
			return
		}

		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	//关闭连接
	defer cli.Close()
}
