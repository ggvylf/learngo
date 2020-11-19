package es

import "fmt"

func Init(addr string) {

	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println("connect to es failed,err=", err)
		return
	}
}
