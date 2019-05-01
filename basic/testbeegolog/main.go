package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {

	//初始化配置
	config := make(map[string]interface{})

	//配置参数
	config["filename"] = "./testbeego.log"
	config["level"] = logs.LevelDebug

	//序列化config
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	//初始化logger
	logs.SetLogger(logs.AdapterFile, string(configStr))

	//生成不同级别的日志
	logs.Debug("this is a debug, my name is %s", "stu01")
	logs.Trace("this is a trace, my name is %s", "stu02")
	logs.Warn("this is a warn, my name is %s", "stu03")
	logs.Info("this is a info, my name is %s", "stu04")

}
