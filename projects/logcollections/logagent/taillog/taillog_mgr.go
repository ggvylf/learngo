package taillog

import (
	"github.com/ggvylf/learngo/projects/logcollections/logagent/etcd"
)
var  (
	tskMgr *tailLogMgr
)


type tailLogMgr sturct {
	logEntryConf []*etcd.LogEntry
	tskMap map[string]*TailTask
	newConfChan chan *etcd.LogEntry
}



func Init(logEntryConf []*etcd.LogEntry)  {
	tskMgr=&tailLogMgr {
		logEntryConf:logEntryConf,
		tskMap: make(map[string]*TailTask,16),
		newConfChan:make(chan []*etcd.LogEntry),
	}


	for _, logEntry := range logEntryConf {

	tailObj:=NewTailTask(logEntry.Path, logEntry.Topic)
	
	//日志文件的路径做key,任务的对象做value
	key:=fmt.Sprintf("%s_%s",logEntry.Path,logEntry.Topic)
	tskMgr.tskMap[key]=tailObj
	}


	go tskMgr.run()

}


func (t *tailLogMgr) run() {
	for {
		select{
		case: newConf:=< -t.newConfChan:
			//增加新配置
			for _,conf:=range newConf {
				key:=fmt.Sprintf("%s_%s",conf.Path,conf.Topic)
				_,ok:=t.tskMap[key]
				if ok {
					continue
				}else {
						tailObj:=NewTailTask(logEntry.Path, logEntry.Topic)
						t.tskMap[key]=tailObj
				}

			}
			//删掉老配置
			//从新老配置列表中比对配置
			for _,c1:=range t.logEntry {
				isDelete=true
				for _,c2:=range newConf {
					if c2.Path == c1.Path && c2.Topic==c1.Topic
					isDelete=false
					continue
				}
				//删除无用的老配置
				if isDelete {
					key:=fmt.Sprintf("%s_%s",c1.Path,c1.Topic)
					t.tskMgr.cancelFunc()
				}
			}
			fmt.Println("新配置增加了")
		default:
			time.Sleep(time.Second)
		}

		

	}
}


