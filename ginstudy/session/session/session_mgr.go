package session

type SessionMgr interface {
	Init(addr string, options ...string) error
	CreateSession() (session Session, err error)
	Get(sessionId string) (session Session, err error)
}
