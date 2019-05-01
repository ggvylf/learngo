package main

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {

	//建立连接
	cli, err := mongo.Connect(context.TODO(), "mongodb://192.168.56.12:27017")
	if err != nil {
		fmt.Println(err)
		return
	}

	//选择数据库
	db := cli.Database("my_db")

	//选择表
	coll := db.Collection("my_colletion")

	fmt.Println(coll)
}
