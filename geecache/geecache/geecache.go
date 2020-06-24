package geecache

import (
	"fmt"
	"log"
	"sync"

	pb "github.com/ggvylf/learngo/geecache/geecache/geecachepb"
	"github.com/ggvylf/learngo/geecache/geecache/singleflight"
)

var (
	mu sync.RWMutex
	//用来存放多个group
	groups = make(map[string]*Group)
)

//Getter接口
type Getter interface {
	Get(key string) ([]byte, error)
}

//定义回调函数类型
type GetterFunc func(key string) ([]byte, error)

//GeeterFunc类型实现了Getter接口的Getf方法
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

//Group类型
//缓存的命名空间，每个Group有唯一的name，
type Group struct {
	name string
	//回调函数
	getter Getter
	//并发缓存
	mainCache cache
	peers     PeerPicker

	//通过loader保证key纸杯调用一次
	loader    *singleflight.Group
}

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	//回调函数如果为nil，就panic
	if getter == nil {
		panic("nil Getter")
	}
	//加锁，保证group的实例唯一的
	mu.Lock()
	defer mu.Unlock()

	//创建group实例
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
		loader:    &singleflight.Group{},
	}

	//把创建的group添加到groups中
	groups[name] = g
	return g
}

//从groups里获取group，返回nil表示group不存在
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()

	return g
}

//从group中获取key的value
//规则：
//流程1. 从mainCache查找，如果存在则返回缓存值
//流程3.  缓存不存在，则调用load方法，load调动GetLocally或者是getFromPeer方法
//getLocally调用用户提供的回调函数 g.getter.Get()来获取源数据，并把通过popluateCache方法把源数据添加到mainCache中
func (g *Group) Get(key string) (ByteView, error) {
	//key为空，返回一个空的ByteView结构体
	if key == "" {
		return ByteView{}, fmt.Errorf("key is requreid")

	}

	//从mainCache中获取对应key的value
	if v, ok := g.mainCache.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}

	//mainCache中没有，开始到其他地方获取key
	return g.load(key)
}

//从本地获取
func (g *Group) getLocally(key string) (ByteView, error) {

	//调用回调函数获取源数据
	bytes, err := g.getter.Get(key)

	//获取不到key，返回空struct
	if err != nil {
		return ByteView{}, err
	}

	//拷贝value
	value := ByteView{b: cloneBytes(bytes)}

	//把kv添加到mainCache中，并且返回value
	g.populateCache(key, value)
	return value, nil

}

//添加到group实例的mainCache中
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}

//把实现了PeerPicker接口的HTTPPool注入到Group中
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}

//从groups中获取key的值
func (g *Group) load(key string) (value ByteView, err error) {

	//把原来的load逻辑，通过g.loader.DO抱起来，确保了并发场景下针对相同的key，load只调用1次
	viewi, err := g.loader.Do(key, func() (interface{}, error) {

		//不是本机节点，远程获取
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok {
				if value, err := g.getFromPeer(peer, key); err == nil {
					return value, nil
				}
				log.Println("[GeeCache] Failed to get from peer", err)
			}
		}

		//本机节点，调用本地方法
		return g.getLocally(key)
	})

	if err != nil {
		return viewi.(ByteView), nil
	}

	return

}

//使用实现了PeerGetter接口的httpGetter从远程节点访问。获取缓存
func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	req := &pb.Request{
		Group: g.name,
		key:   key,
	}

	res := &pb.Response{}
	err := peer.Get(req, res)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: res.Value}, nil

}
