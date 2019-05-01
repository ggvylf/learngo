package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

//时间点
type TimePoint struct {
	StartTime int64 `bson:"starttime"`
	EndTime   int64 `bson:"endtime"`
}

//日志结构
type LogRecord struct {
	JobName   string    `bson:"jobname"`
	Command   string    `bson:"command`
	Err       string    `bson:"err"`
	Content   string    `bson:"content"`
	TimePoint TimePoint `bson:"timepoint"`
}

//查询条件结构体
type FindByJobName struct {
	JobName string `bson:"jobname"`
}

func main() {

	opts := &options.ClientOptions{}
	opts.SetConnectTimeout(5 * time.Second)

	cli, err := mongo.Connect(context.TODO(), "mongodb://192.168.56.11:27017")
	if err != nil {
		fmt.Println(err)
	}

	db := cli.Database("cron")
	coll := db.Collection("log")

	//创建查询条件
	filter := &FindByJobName{
		JobName: "job10",
	}

	//查询
	cursor, err := coll.Find(context.TODO(), filter, options.Find().SetSkip(0), options.Find().SetLimit(10))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.TODO())

	//遍历结果集
	for cursor.Next(context.TODO()) {
		record := &LogRecord{}

		//反序列化结果
		err := cursor.Decode(record)
		if err != nil {
			fmt.Println(err)
			return
		}

		//打印结果
		fmt.Println(*record)
	}

}
