package taillog

import (
	"github.com/ggvylf/learngo/projects/logcollections/logagent/etcd"
)

type tailMgr sturct {
	logEntryConf []*etcd.LogEntry
}