package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

//自定义函数类型Hash，方便自定义或者是替换别的hash函数，默认使用crc32.ChecksumIEEE算法
type Hash func(data []byte) uint32

//一致性hash算法的数据借口
type Map struct {
	//自定义Hash函数
	hash Hash

	//虚拟节点的倍数
	replicas int

	//hash环
	keys []int

	//虚拟节点跟真实节点的映射表，key是虚拟节点的hash值，value是真是节点的名称
	hashMap map[int]string
}

//创建一个Map实例，允许自定义虚拟节点倍数和Hash函数
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}

	//自定义函数为空
	if m.hash == nil {
		//默认的hash算法
		m.hash = crc32.ChecksumIEEE
	}

	return m

}

//添加真实节点
func (m *Map) Add(keys ...string) {
	//对每个真实的节点，创建虚拟节点
	for _, key := range keys {

		for i := 0; i < m.replicas; i++ {
			//创建虚拟节点，虚拟节点的名字=真实节点节个数+真实节点key
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))

			//把虚拟节点添加到hash环上
			m.keys = append(m.keys, hash)

			//增加虚拟节点和真实节点的映射关系
			m.hashMap[hash] = key
		}

		//对环上的hash值排序
		sort.Ints(m.keys)
	}
}

//选择真实节点
func (m *Map) Get(key string) string {
	//hash环长度为0，返回空
	if len(m.keys) == 0 {
		return ""
	}

	//计算key的hash值
	hash := int(m.hash([]byte(key)))

	//顺时针获取第一个虚拟节点的index。
	//如果idx==len(m.keys)，应该选择m.keys[0]
	idx := sort.Search(len(m.keys), func(i int) bool {

		//节点映射map表中值大于key的hash，就是虚拟节点的的hash
		return m.keys[i] >= hash
	})

	//通过节点映射表,找到虚拟节点对应的真实节点
	//因为hash是个一个环状结构，所以通过取余的方式来处理。
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
