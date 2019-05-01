package  main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/mongo"
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

func main() {

	opts := &options.ClientOptions{}
	opts.SetConnectTimeout(5 * time.Second)

	cli, err := mongo.Connect(context.TODO(), "mongodb://192.168.56.11:27017")
	if err != nil {
		fmt.Println(err)
	}

	db := cli.Database("cron")
	coll := db.Collection("log")

	//插入单条记录
	record := &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	result, err := coll.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println(err)
		return
	}
	docId := result.InsertedID
	fmt.Println(docId)


}
