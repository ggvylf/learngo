package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {

	var ()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.56.11:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	kv := clientv3.NewKV(cli)

	//创建op
	putOp := clientv3.OpPut("/hehe", "haha")

	//执行op
	opResp, err := kv.Do(context.TODO(), putOp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入的revision：", opResp.Put().Header.Revision)

	//创建getop
	getOp := clientv3.OpGet("/hehe")
	opResp, err = kv.Do(context.TODO(), getOp)
	fmt.Println("读取数据的revison：", opResp.Get().Kvs[0].ModRevision)
	fmt.Println("读取数据的value：", string(opResp.Get().Kvs[0].Value))

}
