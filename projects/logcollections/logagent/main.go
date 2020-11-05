package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ggvylf/learngo/projects/logcollections/logagent/conf"
	"github.com/ggvylf/learngo/projects/logcollections/logagent/kafka"
	"github.com/ggvylf/learngo/projects/logcollections/logagent/taillog"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func run() {

	//读取日志
	for {
		select {
		//从tailObj.Lines中读取数据
		case line := <-taillog.ReadChan():
			//发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)

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
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed,err=%v\n", err)
		return
	}

	//初始化taillog
	err = taillog.Init(cfg.TailConf.FileName)
	if err != nil {
		fmt.Printf("init tail file failed,err=%v\n", err)
		return
	}

	run()
}
