package main

import "testing"

func TestStore(t *testing.T) {
	m01 := &Monster{"tom", 18, "eat"}
	result := m01.Store()
	if !result {
		t.Logf("函数TestStore()测试失败，期望值%v,实际值%v", true, result)
	} else {
		t.Logf("函数TestStore()测试成功")
	}

}

func TestReStore(t *testing.T) {
	m01 := &Monster{"tom", 18, "eat"}
	result := m01.ReStore()
	if !result {
		t.Logf("函数TestReStore()测试失败，期望值%v,实际值%v", true, result)
	} else {
		t.Logf("函数TestReStore()测试成功")
	}

}
