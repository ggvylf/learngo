package main

import (
	"fmt"
	"testing"
)

//go test -v *.go

// 覆盖率
//go test -v -cover *.go

// // 单元测试
func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {
		ret := Square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d,the expected is %d,the actual %d", inputs[i], expected[i], ret)
		}
	}

}

func TestError(t *testing.T) {
	fmt.Println("test Error start")

	//Fail,Error 该测试失败，该测试继续，其他测试继续执行
	//FailNow,Fatal 该测试失败 该测试终止，其他测试继续执行
	t.Error("error")
	fmt.Println("test Error stop")
}

func TestFatal(t *testing.T) {
	fmt.Println("test Fatal start")

	//Fail,Error 该测试失败，该测试继续，其他测试继续执行
	//FailNow,Fatal 该测试失败 该测试终止，其他测试继续执行
	// t.Fatal("error")
	// 这里不会执行
	fmt.Println("test Fatal stop")
}
