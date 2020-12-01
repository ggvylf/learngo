package session

import (
	"errors"
	"sync"
)

type MemorySesion struct {
	sessionid string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

//工厂函数
func NewMemorySession(id string) *MemorySesion {
	return &MemorySesion{
		sessionid: id,
		data:      make(map[string]interface{}, 16),
	}

}

//接口方法的实现
func (m *MemorySesion) Set(key string, value interface{}) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	m.data[key] = value
	return
}

func (m *MemorySesion) Get(key string) (value interface{}, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	value, ok := m.data[key]

	if !ok {
		err = errors.New("key not exists")
		return
	}
	return
}

func (m *MemorySesion) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	delete(m.data, key)
	return
}

func (m *MemorySesion) Save(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	return
}
