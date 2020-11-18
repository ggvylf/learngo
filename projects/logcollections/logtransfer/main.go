package main

import (
	"fmt"
	"os"

	"github.com/ggvylf/learngo/projects/logcollections/logtransfer/config"
	"gopkg.in/ini.v1"
)

func main() {

	//加载配置文件
	var cfg config.LogTransfer
	err := ini.MapTo(&cfg, "./config/config.ini")
	if err != nil {
		fmt.Println("init config failed,err=", err)
		os.Exit(0)
	}
	fmt.Println(cfg)

	//初始化kafka
	//初始化es
	//从kafka中读取数据
	//写入到es中
}
