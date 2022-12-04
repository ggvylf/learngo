package main

import (
	"testing"
)

// 性能测试

// go test -bench . -v *.go
// go test -v -bench=. -benchmem *.go
func BenchmarkSquare(b *testing.B) {
	inputs := [...]int{1, 2, 3}

	// 重置计时器
	b.ResetTimer()

	//benchmark
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(inputs); i++ {
			Square(inputs[i])
		}
	}
	//停止计时器
	b.StopTimer()

}
