package etcd

import (
	"fmt"
	"time"
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	client *clientv3.Client
)

type LogEntry struct {
	Path string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addr string) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string(addr),
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect etcd failed,err=", err)
		return
	}
	return

}

//从etcd中获取文件和topic列表
func GetConf(key string,timeout int) (logEntryConf []*LogEntry,err error)  {
	ctx,cancel=context.WithTimeout(context context.Background(),timeout)
	resp,err:=client.Get(ctx,key)
	cancel()
	if err!=nil {
		fmt.Println("get key from etcd failed,err=",err)
		return
	}

	for _,ev:=range resp.Kvs {
		err:=json.Unmarshal(ev.Value,&logEntryConf)
		if err!=nil {
			fmt.Println("json unmarshal failed，err=",err)
		}
	}
}