package main

import (
	"fmt"
	"os"
	"time"
)

var msgCh = make(chan string, 10)

// 计算开始结束时间
func timer(inner func(), taskname string) {
	start := time.Now()
	inner()
	end := time.Since(start).Microseconds()
	msg := fmt.Sprintf("%s time used %d ms\n", taskname, end)
	msgCh <- msg
}

// tastA
func taskA() {
	time.Sleep(500 * time.Millisecond)
}

// taskB
func taskB() {
	time.Sleep(4000 * time.Millisecond)
}

func main() {

	timer(taskA, "taskA")
	timer(taskB, "taskB")

	//读取msgCh超时就退出

	for {
		select {
		case ret := <-msgCh:
			fmt.Println(ret)

		// 超时控制
		case <-time.After(1000 * time.Microsecond):
			fmt.Println("wait msgCh timeout")
			os.Exit(0)

		}
	}

}
