package main

import (
	"fmt"
	"os"

	"github.com/ggvylf/learngo/projects/logcollections/logtransfer/config"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(config.LogTransfer)
)

func main() {

	//加载配置文件

	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Println("init config failed,err=", err)
		os.Exit(0)
	}

	//初始化kafka
	//初始化es
	//从kafka中读取数据
	//写入到es中
}
