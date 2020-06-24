package geecache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Fatal("callback failed")
	}
}

func TestGet(t *testing.T) {

	//回调函数次数，某个key对应调用回调的次数
	loadCounts := make(map[string]int, len(db))

	//初始化group
	gee := NewGroup("scores", 2<<10, GetterFunc(
		//回调函数
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			//查看key是否再db中
			if v, ok := db[key]; ok {
				//是否调用回调函数，如果这个key的调用次数不存在，那么就把回调函数的次数设置为0
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}
				//如果有调用 次数+1
				loadCounts[key]++

				//返回key对应的v
				return []byte(v), nil
			}
			//key不在db中
			return nil, fmt.Errorf("%s not exist", key)
		}))

	for k, v := range db {
		//缓存为空
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatal("failed to get value of Tom")
		}
		//回调函数>1 表示多次调用回调函数，没有缓存
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}

	//key不存在
	if view, err := gee.Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty, but %s got", view)
	}
}

func TestGetGroup(t *testing.T) {
	groupName := "scores"
	NewGroup(groupName, 2<<10, GetterFunc(
		func(key string) (bytes []byte, err error) { return }))
	if group := GetGroup(groupName); group == nil || group.name != groupName {
		t.Fatalf("group %s not exist", groupName)
	}

	if group := GetGroup(groupName + "111"); group != nil {
		t.Fatalf("expect nil, but %s got", group.name)
	}
}
