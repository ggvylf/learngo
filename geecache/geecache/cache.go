package geecache

import (
	"sync"

	"github.com/ggvylf/learngo/geecache/geecache/lru"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

//cache的add方法
func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()

	//延迟初始化,为nil才初始化
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}

	//把key和value添加到c的链表中
	c.lru.Add(key, value)
}

//cache的get方法
func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	//缓存为空
	if c.lru == nil {
		return
	}

	//key存在
	if v, ok := c.lru.Get(key); ok {
		//断言，制定返回ByteView类型
		return v.(ByteView), ok
	}

	return
}
