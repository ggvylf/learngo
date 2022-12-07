// GO111MODULE="auto" go build *.go
// 查看cpu cpu go tool pprof main cpu.prof
// 查看内存 GO111MODULE="auto" go build *.go
// 查看goroutine

package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

const (
	col = 10000
	row = 10000
)

func fillMatrix(m *[row][col]int) {
	// 生成随机数
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 填充矩阵
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(10000)
		}
	}
}

// 计算矩阵内数据的sum
func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

func main() {

	// cpu相关
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Sprintln("could not create CPU profile:", err)
	}

	// 开始监控cpu
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Sprintln("could not start CPU profile: ", err)
	}
	// 结束监控cpu
	defer pprof.StopCPUProfile()

	// 主要执行逻辑
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	// mem相关
	fm, err := os.Create("mem.prof")
	if err != nil {
		fmt.Sprintln("could not create Mem profile:", err)
	}

	// 可以查看GC前后的内存占用情况
	// runtime.GC()
	if err := pprof.WriteHeapProfile(fm); err != nil {
		fmt.Sprintln("could not write memory profile: ", err)
	}
	fm.Close()

	//goroutine

	fg, err := os.Create("goroutine.prof")
	if err != nil {
		fmt.Sprintln("could not create groutine profile: ", err)
	}

	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		fmt.Sprintln("could not write groutine profile: ")
	} else {
		gProf.WriteTo(fg, 0)
	}
	fg.Close()
}
