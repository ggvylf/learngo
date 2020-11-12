package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

//一个日志收集的任务的实例
type TailTask struct {
	path     string `json:"path"`
	topic    string `json:"topic`
	instance *tail.Tail
}

func NewTailTask(path,topic string) (tailObj *tail.TailTask) {
	tailObj=&TailTask {
		path:path,
		topic:topic
	}
	tailObj.Init(tailObj.path)
	return
}


func (t *TailTask) Init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed,err=", err)

	}



}

func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}
