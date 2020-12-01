package session

import (
	"sync"

	"github.com/google/uuid"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

//工厂函数
func NewMemorySessionMap() *MemorySessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}

}

func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	return
}
