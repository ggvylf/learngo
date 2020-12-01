package session

//定义接口类型
type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}
