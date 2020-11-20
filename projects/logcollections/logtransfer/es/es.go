package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic"
)

var (
	client *elastic.Client
)

func Init(addr string) (err error) {
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
	return

}

//发送数据到es
func SendToES(indexStr string, data interface{}) {

	//链式操作，需要在对应的func中返回操作的对象
	// 创建的document为 /user/go/
	put1, err := client.Index().Index(indexStr).
		Type("go").
		BodyJson(p1).
		Do(context.Background())

	if err != nil {
		fmt.Println("put key to index failed,err=", err)
		return
	}
	return
}
