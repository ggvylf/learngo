package session

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/gomodule/redigo/redis"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
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
	r.flag = SessionFlagModify

	return

}

//存储数据到redis
func (r *RedisSession) Save(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	if r.flag != SessionFlagModify {
		return
	}

	//把内存数据序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}

	//获取redis连接
	conn := r.pool.Get()
	//保存kv到redis中
	_, err = conn.Do("SET", r.sessionId, string(data))
	if err != nil {
		return
	}
	//保存后修改状态
	r.flag = SessionFlagNone

	return

}

func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	//先判断内存中是否有
	if r.flag == SessionFlagNone {
		result, ok := r.sessionMap[key]
		if !ok {
			err := errors.New("key not in memory")
			return
		}

	}
	return

}
func (r *RedisSession) LoadFromRedis(key string) (err error) {
	conn := r.pool.Get()
	result, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}

	//转字符串
	data, err := redis.String(result, err)
	if err != nil {
		return
	}

	//把data重新写入内存
	err = json.Unmarshal([]byte(data), r.sessionMap)

	if err != nil {
		return
	}
	return
}

func (r *RedisSession) Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	r.flag = SessionFlagModify
	delete(r.sessionMap, key)

	return
}
