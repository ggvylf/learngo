package singleflight

import (
	"sync"
)

//请求结构体， 表示正在进行中或者是已结束的请求
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

//请求组，管理不同key的call
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

//针对同一个key,无论Do()被调用多少次，fn()都只会被调用一次
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {

	//加锁，保护g.m不被并发读写
	g.mu.Lock()

	if g.m == nil {
		g.m = make(map[string]*call)
	}

	//g.m[key]中有数据，表示当前有请求正在请求
	if c, ok := g.m[key]; ok {
		//解锁g.m
		g.mu.Unlock()

		//阻塞c
		c.wg.Wait()

		//返回结果
		return c.val, c.err
	}

	//c不存在，初始化新的请求，延迟初始化
	c := new(call)

	//请求前加锁
	c.wg.Add(1)

	//把c写入g.m，表示key已经有对应的请求在处理
	g.m[key] = c

	//解锁g.m
	g.mu.Unlock()

	//调用fn(),获取返回值
	c.val, c.err = fn()

	//请求完成，解锁c
	c.wg.Done()

	//加锁 保护gm
	g.mu.Lock()
	//从g.m中删除key
	delete(g.m, key)

	//解锁g.m
	g.mu.Unlock()

	return c.val, c.err

}
