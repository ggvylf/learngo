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

//工厂模式，新建一个task实例
func NewTailTask(path,topic string) (tailObj *tail.TailTask) {
	tailObj=&TailTask {
		path:path,
		topic:topic
	}
	tailObj.init(tailObj.path)
	return
}


func (t *TailTask) init() {
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


	//发送日志到kafka
	go t.run()


}

func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}


// 发送到kafka
func (t *TailTask) run() {
	
	for {
		select {
		case line := <-t.instance.Lines:
			kafka.SentToKafka(logEntry.topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}