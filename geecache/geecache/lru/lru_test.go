package lru

import (
	"reflect"
	"testing"
)

type String string

//String类型实现了Value类型的方法
func (d String) Len() int {
	return len(d)
}

//测试Get
func TestGet(t *testing.T) {
	lru := New(int64(0), nil)

	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failee")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}

}

//测试内存超过设定值
func TestRemovedoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"

	cap := len(k1 + k2 + v1 + v2)
	//初始化的cap只包含kv1和kv2
	lru := New(int64(cap), nil)

	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	//插入kv3后，超出cap，key1会被清除
	lru.Add(k3, String(v3))

	//如果key1存在或者是lru的长度不等于2，则表示旧的key没有被清除
	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}


}

//测试回调
func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	//回调函数，把keyappend到keys中
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(int64(10), callback)
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("k2"))
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4"))

	//期望的结果
	expect := []string{"key1", "k2"}


	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}

//测试Add
func TestAdd(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key", String("1"))
	lru.Add("key", String("111"))

	if lru.nbytes != int64(len("key")+len("111")) {
		t.Fatal("expected 6 but got", lru.nbytes)
	}
}
