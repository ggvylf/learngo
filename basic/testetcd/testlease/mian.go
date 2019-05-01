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

		fmt.Println("链接服务器失败，错误是:", err)
		return
	}

	fmt.Println("链接服务器成功")
	//关闭连接
	defer cli.Close()

	//创建一个租约
	lease := clientv3.NewLease(cli)

	//指定租约时间为10s
	leaseGrantResp, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println(err)
	}

	//获取租约id
	leaseId := leaseGrantResp.ID

	//自动续租
	keepRespChan, err := lease.KeepAlive(context.TODO(), leaseId)
	if err != nil {
		fmt.Println(err)
		return
	}
	// //ctx超时模拟
	// //此时ctx续租时间5秒，租约10s，共计15s
	// ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	// keepRespChan, err := lease.KeepAlive(ctx), leaseId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//查看自动续租chan的状态
	go func() {
		for {
			select {
			case keepResp := <-keepRespChan:
				//网络连接失败或者是context超时
				if keepRespChan == nil {
					fmt.Println("租约失效")
					goto END
				} else {
					//每秒续租一次
					fmt.Println("收到自动续租应答", keepResp.ID)
				}
			}
		END:
		}
	}()

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
			fmt.Println("租约超时")
		} else {
			fmt.Println("租约仍有效")
		}

		time.Sleep(time.Second)

	}
}
