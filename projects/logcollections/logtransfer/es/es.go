package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic"
)


var (
	client *elastic.Client
	ch chan *LogData
)

type LogData struct {
	Topic string `json:"topic"`
	Data string `json:"data"`
}

func Init(addr string,chansize int,chancount int) (err error) {
	//如果es地址不带http://，填充
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}

	//创建一个客户端
	client, err = elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		fmt.Println("connect to es failed,err=", err)
		return
	}
	//连接建立厚就可以往chan中写入数据了
	ch=make(chan *LogData,chansize)

	for i:=0;i<chancount;i++ {
		go SendToESChan()
	}


	return

}

//发送数据到es的chan中
func SendToESChan(msg *LogData) {

	ch <- msg
}


//发送数据到es
func SendToES() {

	for (
		select {
		case msg :=<- ch : 
				put1, err := client.Index().Index(msg.Topic).
				Type("go").
				BodyJson(msg).
				Do(context.Background())

			if err != nil {
				fmt.Println("put key to index failed,err=", err)
				continue
			}

default:
	time.Sleep(time.Second)
		}
	)


}

