package consistenthash

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	//新建一个3倍虚拟节点的hasn环，真实节点个数为3个，总计9个节点
	//
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	//添加真实节点到hash环
	//hash环上的节点：
	// 02/12/22、04/14/24、06/16/26
	hash.Add("6", "4", "2")

	//数字和对应的hash节点
	testCases := map[string]string{

		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	// 增加一个真实节点8
	//hash环上的节点：
	// 02/12/22、04/14/24、06/16/26 08/18/28
	hash.Add("8")
	fmt.Println(hash)

	// 27对应的真实节点节点应该是8
	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

}
