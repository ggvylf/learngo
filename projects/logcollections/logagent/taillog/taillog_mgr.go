package taillog

import (
	"github.com/ggvylf/learngo/projects/logcollections/logagent/etcd"
)



var  (
	tskMgr *tailLogMgr
)


type tailLogMgr sturct {
	logEntryConf []*etcd.LogEntry
	// tskMap map[string]TailTask
}



func Init(logEntryConf []*etcd.LogEntry)  {
	tskMgr=&tailLogMgr {
		logEntryConf:logEntryConf,
	}


	//循环收集日志发送到kafka
	for _, logEntry := range logEntryConf {
	NewTailTask(logEntry.Path, logEntry.Topic)
	}

}


