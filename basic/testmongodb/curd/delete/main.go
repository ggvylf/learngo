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

//删除操作符结构体
//{"$lt":timestamp}
type TimeBeforeCond struct {
	Before int64 `bson:"$lt"`
}

//删除条件结构体
//delete{"timePoint.startTime":{"$lt":timestamp}}
type DeleteCond struct {
	BeforeCond TimeBeforeCond `bson:"timepoint.starttime"`
}

func main() {

	opts := &options.ClientOptions{}
	opts.SetConnectTimeout(5 * time.Second)

	cli, err := mongo.Connect(context.TODO(), "mongodb://192.168.56.11:27017", opts)
	if err != nil {
		fmt.Println(err)
	}

	db := cli.Database("cron")
	coll := db.Collection("log")

	//delete({"timepoint.startime":{"$lt":当前时间}})
	deletecond := &DeleteCond{BeforeCond: TimeBeforeCond{Before: time.Now().Unix()}}

	delresult, err := coll.DeleteMany(context.TODO(), deletecond)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(delresult.DeletedCount)
}
