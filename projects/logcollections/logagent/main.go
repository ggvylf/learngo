package logagent

import (
	"gopkg.in/ini.v1"
	"fmt"

	"github.com/ggvylf/learngo/projects/logcollections/logagent/kafka"
	"github.com/ggvylf/learngo/projects/logcollections/logagent/taillog"
)

func run() {
	//读取配置文件
	configfile:="./config.ini"
	cfg,err:=ini.Load(configfile)
	if err!=nil {
		fmt.Println("load config file failed,err=",err)
		os.Exit(1)
	}


	




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

	//初始化kafka
	kafkaddr := []string{"127.0.0.1:9092"}
	err := kafka.Init(kafkaddr)
	if err != nil {
		fmt.Printf("init kafka failed,err=%v\n", err)
		return
	}

	//初始化taillog
	filename := "./a.log"
	err = taillog.Init(filename)
	if err != nil {
		fmt.Printf("init tail file failed,err=%v\n", err)
		return
	}

	run()
}
