package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Pool *redis.Pool
var wg sync.WaitGroup

func init() {
	redispwd := redis.DialPassword("deepin")
	Pool = &redis.Pool{
		MaxIdle:     20,
		MaxActive:   0,
		IdleTimeout: time.Second * 30,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379", redispwd)
		},
	}
}

func get(Pool *redis.Pool, i int) {
	conn := Pool.Get()
	defer conn.Close()
	rep2, err := conn.Do("get", "name")
	res2, _ := redis.String(rep2, err)
	fmt.Println(i, res2)

	wg.Done()

}

func main() {
	defer Pool.Close()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go get(Pool, i)
	}
	wg.Wait()

}
