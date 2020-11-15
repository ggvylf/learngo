package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age`
	Married bool   `json:"married"`
}

func main() {
	//初始化连接
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println("connect to es failed,err=", err)
		return
	}

	p1 := Person{
		Name:    "tom",
		Age:     18,
		Married: false,
	}

	//链式操作，需要在对应的func中返回操作的对象
	put1, err := client.Index().Index("user").BodyJson(p1).Do(context.Background())

	if err != nil {
		fmt.Println("put key to index failed,err=", err)
		return
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

}
