package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func putEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.56.11:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()

	if err != nil {
		fmt.Println("connect failed, err:", err)
		panic("connect failed")
	}

	fmt.Println("connect succ")

	//设置1秒超时，访问etcd有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//操作etcd
	_, err = cli.Put(ctx, "/mytestetcd/k1", "v1")
	//操作完毕，取消etcd

	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
	fmt.Println("put ok!")
	cancel()

}

func getEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.56.11:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
	if err != nil {
		fmt.Println("connect failed, err:", err)
		panic("connect failed")
	}

	fmt.Println("connect succ")

	//取值，设置超时为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/mytestetcd/k1")
	cancel()

	if err != nil {
		fmt.Println("get etcd key failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

}

func main() {

	putEtcd()
	getEtcd()

}
