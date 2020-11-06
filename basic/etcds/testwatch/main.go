package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
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

	//模拟kv变化
	go func() {
		for {
			kv.Put(context.TODO(), "/hehe", "haha")
			kv.Delete(context.TODO(), "/hehe")
			time.Sleep(1 * time.Second)
		}

	}()

	//先get到/hehe的值
	getResp, err := kv.Get(context.TODO(), "/hehe")
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值", string(getResp.Kvs[0].Value))
	}
	//获取事务id+1作为监听的开始id
	watcherStartRevision := getResp.Header.Revision + 1

	//创建一个watcher
	watcher := clientv3.NewWatcher(cli)

	//启动监听
	fmt.Println("从该rev以后开始监听", watcherStartRevision)

	//主动终止watch，主动取消context
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})

	// watchRespChan := watcher.Watch(context.TODO(), "/hehe", clientv3.WithRev(watcherStartRevision))
	watchRespChan := watcher.Watch(ctx, "/hehe", clientv3.WithRev(watcherStartRevision))

	//遍历chan中的内容
	for watchResp := range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了，Revision：", event.Kv.ModRevision)
			}
		}
	}
}
