package main

import (
	"gopkg.in/ini.v1"
	"fmt"

	"github.com/ggvylf/learngo/projects/logcollections/logagent/kafka"
	"github.com/ggvylf/learngo/projects/logcollections/logagent/taillog"
)

func run() {

	//读取日志
	for {
		select {
		case line := <- taillog.ReadChan() :
			//发送到kafka
			kafuka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		
		}


}

func main() {

	//读取配置文件
	configfile:="./config.ini"
	cfg,err:=ini.Load(configfile)
	if err!=nil {
		fmt.Println("load config file failed,err=",err)
		os.Exit(1)
	}




	//初始化kafka
	kafkaddr := []string{cfg.Section("kafka").Key("address")}
	err := kafka.Init(kafkaddr)
	if err != nil {
		fmt.Printf("init kafka failed,err=%v\n", err)
		return
	}

	//初始化taillog
	err = taillog.Init(cfg.Section("taillog").Key("path"))
	if err != nil {
		fmt.Printf("init tail file failed,err=%v\n", err)
		return
	}

	run()
}
