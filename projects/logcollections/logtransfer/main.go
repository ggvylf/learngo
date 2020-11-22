package main

import (
	"fmt"
	"os"

	"github.com/ggvylf/learngo/projects/logcollections/logagent/kafka"
	"github.com/ggvylf/learngo/projects/logcollections/logtransfer/config"
	"github.com/ggvylf/learngo/projects/logcollections/logtransfer/es"
	"gopkg.in/ini.v1"
)

func main() {

	//加载配置文件
	var cfg = new(config.LogTransfer)
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Println("init config failed,err=", err)
		os.Exit(0)
	}
	fmt.Println(cfg)

	//初始化es
	err := es.Init(cfg.EsConf.Address, cfg.EsConf.ChanSize, cfg.EsConf.ChanCount)
	if err != nil {
		fmt.Println("init es failed,err=", err)
	}

	//初始化kafka
	err = kafka.Init([]strings{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Println("init kafka failed, err=", err)
		return
	}

	//从kafka中读取数据
	//写入到es中
}
