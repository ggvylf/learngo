package session

import (
	"encoding/json"
	"sync"

	"github.com/gomodule/redigo/redis"
)

const (
	SessionFlagNone = iota
	SeesionFlagModify
)

type RedisSession struct {
	sessionId string
	pool      *redis.Pool

	//session先存在内存中 可以批量写入redis
	sessionMap map[string]interface{}
	//读写锁
	rwlock sync.RWMutex
	//记录内存中的map是否被操作
	flag int
}

//构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	return &RedisSession{
		sessionId:  id,
		sessionMap: make(map[string]interface{}, 16),
		pool:       pool,
		flag:       SessionFlagNone,
	}
}

func (r *RedisSession) Set(key string, value interface{}) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	r.sessionMap[key] = value
	r.flag = SeesionFlagModify

	return

}

//存储数据到redis
func (r *RedisSession) Save(key string) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	if r.flag != SeesionFlagModify {
		return
	}

	//把内存数据序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}

	//获取redis连接
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionId, string(data))
	if err != nil {
		return
	}

	return

}
