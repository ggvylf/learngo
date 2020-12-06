package session

import (
	"errors"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type RedisSessionMgr struct {

	//redis地址
	addr string
	//密码
	passwd string

	//连接池
	pool *redis.Pool

	//锁
	rwlock sync.RWMutex

	//大map
	sessionMap map[string]Session
}

//构造函数
func NewRedisSessionMgr() SessionMgr {
	return &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	//处理多余参数
	if len(options) > 0 {
		r.passwd = options[0]
	}

	//创建连接池
	r.pool = myPool(addr, r.passwd)
	r.addr = r.addr
	return

}

//自定义redis连接池
func myPool(addr, passwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			//判断是否有密码
			if _, err := conn.Do("auth", passwd); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err

		},

		//测试连接
		//上线时要去掉
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err

		},
	}

}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	id, err := uuid.NewRandom()
	if err != nil {
		return
	}

	sessionId := id.String()

	session = NewRedisSession(sessionId, r.pool)

	//把session根据sessinId加入到session的map中
	r.sessionMap[sessionId] = session

	return
}
func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {

	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return

}
