package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	//建立连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("context deadline")
			return
		}

		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect etcd success")
	//关闭连接
	defer cli.Close()

	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, "hehe", "haha")
	cancel()
	if err != nil {
		fmt.Println("put key to etcd failed,err=", err)
		return
	}

	//get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.GET("hehe")
	cancel()
	if err != nil {
		fmt.Println("get key from etcd failed,err=", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("key=%v,value=%v\n", ev.Key, ev.Value)
	}

	// //使用clientv3.WithPrevKV()参数，获取上一次的值
	// putResp, err := cli.Put(context.TODO(), "/hehe", "lala", clientv3.WithPrevKV())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(putResp.Header)

	// //如果前一个pv不为空，说明做出了修改，打印上一次修改的value
	// if putResp.PrevKv!=nil {
	// fmt.Println(string(putResp.PrevKv.Value))

	// }

	// getResp, err := cli.Get(context.TODO(), "/", clientv3.WithPrefix())
	// if err != nil {
	// 	fmt.Println(err)

	// }

	// for _, kv := range getResp.Kvs {
	// 	fmt.Println(string(kv.Key), string(kv.Value))
	// }

	// delResp, err := cli.Delete(context.TODO(), "/hehe", clientv3.WithPrevKV())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if len(delResp.PrevKvs) != 0 {
	// 	for _, kv := range delResp.PrevKvs {
	// 		fmt.Println("deleted", string(kv.Key), string(kv.Value))
	// 	}
	// }

}
