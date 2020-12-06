package session

import (
	"fmt"
)

var (
	sessionMgr SessionMgr
)

//选择使用版本

func Init(provider string, addr string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Errorf("not support")
		return
	}

	err = sessionMgr.Init(addr, options...)
	return

}
