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

	//创建一个租约
	lease := clientv3.NewLease(cli)

	//指定租约时间为10s
	leaseGrantResp, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println(err)
	}

	//获取租约id
	leaseId := leaseGrantResp.ID

	//创建kv子集
	kv := clientv3.NewKV(cli)

	//写入etcd，指定租约
	putResp, err := kv.Put(context.TODO(), "/hehe", "haha", clientv3.WithLease(leaseId))
	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(putResp.Header)

	//获取key的时候指定超时
	for {
		getResp, err := kv.Get(context.TODO(), "/hehe", clientv3.WithLease(leaseId))
		if err != nil {
			fmt.Println(err)
		}

		if getResp.Count == 0 {
			fmt.Println("lease is timeout")
		} else {
			fmt.Println("leaseing ")
		}

		time.Sleep(time.Second)

	}
}
