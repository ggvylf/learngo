package geecache

import (
	"fmt"
	"log"
	"sync"
)

//接口
type Getter interface {
	Get(key string) ([]byte, error)
}

//定义函数类型
type GetterFunc func(key string) ([]byte, error)

//实现了Getter接口
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nill Getter")
	}
	mu.Lock()
	defer mu.Unlock()

	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}

	groups[name] = g
	return g
}

//从groups获取group
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()

	return g
}

//从group中获取key
func (g *Group) Get(key string) (ByteView, error) {
	//key为空，返回一个空的[]byte
	if key == "" {
		return ByteView{}, fmt.Errorf("key is requreid")

	}

	//从mainCache中获取数据
	if v, ok := g.mainCache.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}

	//mainCache中没有，开始到其他地方获取key
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	return g.getLocally(key)

}

//从本地获取
func (g *Group) getLocally(key string) (ByteView, error) {

	//调用回调函数
	bytes, err := g.getter.Get(key)
	//获取不到key，返回空sturct
	if err != nil {
		return ByteView{}, err
	}

	value := ByteView{b: cloneBytes(bytes)}

	//把kv添加到mainCache中，并且返回value
	g.populateCache(key, value)
	return value, nil

}

//添加到group实例的mainCache中
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
