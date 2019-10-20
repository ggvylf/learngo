package main

import (
	"testing"
)

func TestMyDecode(t *testing.T) {
	srcfile := "./info.json"
	p := MyDecode(srcfile)
	if p.Name == "tom" && p.Age == 18 {
		t.Log("单元测试成功")
	} else {
		t.Error("单元测试失败")
	}
}

func TestMyEncode(t *testing.T) {
	destfile := "./a.txt"
	err := MyEncode(destfile)
	if err != nil {
		t.Error("单元测试失败")
	} else {
		t.Log("单元测试成功")
	}
}

func BenchmarkMyDecode(b *testing.B) {
	srcfile := "./info.json"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MyDecode(srcfile)
	}

}

func BenchmarkMyEncode(b *testing.B) {

	destfile := "./a.txt"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MyEncode(destfile)
	}
}
