package main

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

func main() {
	redispwd := redis.DialPassword("deepin")
	Conn, err := redis.Dial("tcp", "127.0.0.1:6379", redispwd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer Conn.Close()
	rep1, err := Conn.Do("set", "age", 18)
	res1, _ := redis.String(rep1, err)
	fmt.Println(res1)

	rep2, err := Conn.Do("get", "name")
	res2, _ := redis.String(rep2, err)
	fmt.Println(res2)

	rep3, err := Conn.Do("get", "age")
	res3, _ := redis.Int(rep3, err)
	fmt.Println(res3)
}
