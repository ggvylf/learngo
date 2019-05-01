package main

import (
	"testing"
)

func TestSub(t *testing.T) {
	res := Sub(1, 2)
	if res != 3 {
		t.Fatalf("Sub(1,2)执行错误,期望值=%v，实际值=%v", 3, res)
	}
	t.Logf("Sub(1,2)执行正确")

}
