// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	//获取当前cpu个数
// 	cpuNum := runtime.NumCPU()
// 	fmt.Println("CPU总个数=", cpuNum)

// 	//忽略单核
// 	if cpuNum != 1 {
// 		runtime.GOMAXPROCS(cpuNum - 1)
// 	}

// }

package main

import (
    "fmt"
    "runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
    //runtime.Gosched()表示把cpu时间片让给其他task，下次某个时候继续回复执行该goroutine
        runtime.Gosched()
        fmt.Println(s)
    }
}

func main() {
    go say("goroutine") //开一个新的Goroutines执行
    say("main")    //当前Goroutines执行
}

