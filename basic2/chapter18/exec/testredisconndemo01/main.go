package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//密码
	redispwd := redis.DialPassword("deepin")
	//数据库
	redisdb := redis.DialDatabase(0)

	//建立链接

	pool := &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379", redispwd, redisdb)
		},
	}

	c := pool.Get()
	defer pool.Close()

	//批量set

	_, err := c.Do("mset", "name1", "tom", "name2", "ann")
	if err != nil {
		fmt.Println(err)
	}

	r, err := redis.Strings(c.Do("mget", "name1", "name2"))
	for _, v := range r {
		fmt.Printf("%v\n", v)
	}

	_, err = c.Do("mset", "name3", "aaa", "name4", "bbb")
	if err != nil {
		fmt.Println(err)
	}

	r, err = redis.Strings(c.Do("mget", "name3", "name4"))
	for _, v := range r {
		fmt.Printf("%v\n", v)
	}
}
