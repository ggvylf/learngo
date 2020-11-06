package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.56.11:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()

	//1. 加锁

	//创建一个租约
	lease := clientv3.NewLease(cli)

	//指定租约时间为5s
	leaseGrantResp, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println(err)
	}

	//获取租约id
	leaseId := leaseGrantResp.ID

	//创建一个自动过期的context
	ctx, cancelFunc := context.WithCancel(context.TODO())

	//确保函数退出后停止自动续租
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)

	//自动续租
	keepRespChan, err := lease.KeepAlive((ctx), leaseId)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	//抢锁
	kv := clientv3.NewKV(cli)

	//创建事务
	txn := kv.Txn(context.TODO())

	//定义事务，key不存在就put，否则就抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("/hehe"), "=", 0)).
		Then(clientv3.OpPut("/hehe", "haha", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/hehe"))

	//提交事务
	txnResp, err := txn.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}

	//判断事务结果
	if !txnResp.Succeeded {
		fmt.Println("锁被占用", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	//2.业务处理

	fmt.Println("业务处理部分")
	time.Sleep(5 * time.Second)

	//3. 释放锁
	//defer来处理

}
