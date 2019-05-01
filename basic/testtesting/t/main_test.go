package main

import (
	"fmt"
	"testing"
)

func testprint1(t *testing.T) {
	result := Print1(2)
	if result != 3 {
		fmt.Println("func test Print1 failed")
		return
	}
}
func testprint2(t *testing.T) {
	result := Print2(2)
	if result != 4 {
		fmt.Println("func test Print2 failed")
		return
	}
}

func TestAll(t *testing.T) {
	t.Run("testprint1", testprint1)
	t.Run("testprint2", testprint2)

}

func TestMain(m *testing.M) {
	fmt.Println("TestMain run first")
	m.Run()
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1(2)
	}
}
