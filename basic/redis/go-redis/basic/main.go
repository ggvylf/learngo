package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client
var ctx = context.Background()

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return
}

func zadd() {

	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 100, Member: "tom"},
		redis.Z{Score: 80, Member: "jerry"},
		redis.Z{Score: 70, Member: "mike"},
	}

	//追加元素到key里
	num, err := redisdb.ZAdd(key, items...).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("nun=", num)

}

func zrange() {
	rst, _ := redisdb.ZRange("rank", 0, 100).Result()
	fmt.Printf("%v\n", rst)
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("连接redis失败")
		return
	}
	fmt.Println("连接redis成功")

	zadd()
	zrange()

}
