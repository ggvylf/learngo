package session

import (
	"errors"
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

	id, err := uuid.NewRandom()
	if err != nil {
		return
	}

	sessionId := id.String()

	session = NewMemorySession(sessionId)

	//把session根据sessinId加入到session的map中
	m.sessionMap[sessionId] = session

	return
}
func (m *MemorySessionMgr) Get(sessionId string) (session Session, err error) {

	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	session, ok := m.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return

}
