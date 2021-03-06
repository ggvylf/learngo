package lru

import "container/list"

type Cache struct {

	//允许使用的最大内存，
	maxBytes int64

	//当前已使用的内容
	nbytes int64

	//双向链表
	ll *list.List

	//缓存对象字典 k是string，v是对应节点的指针
	cache map[string]*list.Element

	//key被删除的回调函数，可以为nil
	OnEvicted func(key string, value Value)
}

//双向链表中节点的数据类型
type entry struct {
	key string
	//占用内存大小
	//值是实现了Value接口的任意类型。只要实现Len()方法即可
	value Value
}

//接口类型，返回值所占用的内存大小
type Value interface {
	Len() int
}

//New()函数
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//查找key
// 1. 先从map中找到对应链表的节点
// 2. 把该节点移动到链表的尾部
func (c *Cache) Get(key string) (value Value, ok bool) {
	//判断缓存链表中，key对应的节点是否存在，如果存在把节点赋值给ele
	if ele, ok := c.cache[key]; ok {
		//把ele移动到对尾
		c.ll.MoveToFront(ele)

		//拿到对应的kv
		//类型断言，从空接口中获取指定类型的值，这里是*entry类型，包含了key和Value
		kv := ele.Value.(*entry)
		//返回kv中的value
		return kv.value, true
	}
	//key不存在，直接返回
	return
}

//移除key
//缓存淘汰，把链表队首的节点移除
func (c *Cache) RemoveOldest() {
	//从ll中获取最后一个元素，如果nil表示链表是空的
	ele := c.ll.Back()

	if ele != nil {
		//从ll中删除ele
		c.ll.Remove(ele)

		//拿到ele的kv
		kv := ele.Value.(*entry)

		//从cache的map中删除kv
		delete(c.cache, kv.key)

		//cache的已使用大小=当前已使用的大小nbytes-(key的长度+value值的长度)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())

		//删除key后，如果回调函数不为nil，调用回调函数 把kv的key和value返回
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}

}

//添加和修改key
func (c *Cache) Add(key string, value Value) {
	//判断key是否存在
	//key存在，只要更新key的value和nbytes
	if ele, ok := c.cache[key]; ok {
		//移动key到队尾
		c.ll.MoveToFront(ele)

		//获取key对应的value
		kv := ele.Value.(*entry)

		//已使用的内存=当前已使用的内存+（新的value的长度-老value的长度）
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())

		//更新对应key的value
		kv.value = value

		//key不存在
	} else {
		//把key写入队尾
		ele := c.ll.PushFront(&entry{key, value})
		//添加映射关系
		c.cache[key] = ele

		//调整已分配大小
		//已分配大小=当前已分配的大小+key的长度+value的长度
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	//清理最少访问的key
	//条件 c.maxbytes不为0 并且nbytes的大小超过maxBytes
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

//返回ll的长度
func (c *Cache) Len() int {
	return c.ll.Len()
}
