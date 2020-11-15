package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ggvylf/learngo/projects/logcollections/logagent/conf"
	"github.com/ggvylf/learngo/projects/logcollections/logagent/etcd"
	"github.com/ggvylf/learngo/projects/logcollections/logagent/kafka"
	"github.com/ggvylf/learngo/projects/logcollections/logagent/taillog"
	"github.com/ggvylf/learngo/projects/logcollections/logagent/utils"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic`
}

func run() {

	//读取日志
	for {
		select {
		//从tailObj.Lines中读取数据
		case line := <-taillog.ReadChan():
			//发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(500 * time.Millisecond)

		}

	}

}

func main() {

	//读取配置文件

	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("map conf failed,err=", err)
		os.Exit(1)
	}

	//初始化kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed,err=%v\n", err)
		return
	}

	//获取本机ip
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	//替换etcd中key的路径
	etcdConfKey := fmt.Sprintf(cfg.EetcdConf.Key, ipStr)

	//初始化etcd
	err = etcd.Init([]string(cfg.EtcdConf.Address), cfg.EtcdConf.Timeout)
	if err != nil {
		fmt.Printf("init etcd failed,err=%v\n", err)
		return
	}

	logentryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fm.Println("etcd.GetConf failed,err=", err)
		return
	}

	//收集日志发往kafka
	taillog.Init(logEntryConf)

	//watch配置文件是否变更
	newConfChan := taillog.newConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan)
	wg.wait()

	run()
}
