package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//创建条件变量
	condvar := false

	//创建一个条件控制器，必须是通过互斥锁来保护
	conder := sync.NewCond(&sync.Mutex{})

	//变量控制部分
	go func() {
		for {

			//修改条件之前要加锁
			conder.L.Lock()
			// 把判断条件设置为true
			condvar = true

			//向conder发送信号
			conder.Signal()
			// conder.Broadcast()
			conder.L.Unlock()

			time.Sleep(1 * time.Second)

			//重置条件，模拟条件变化
			conder.L.Lock()
			condvar = false
			conder.L.Unlock()
			time.Sleep(3 * time.Second)
		}
	}()

	//判断部分
	for {
		//先锁定
		conder.L.Lock()

		for !condvar {
			fmt.Println("不满足控制条件，等待")
			//这里会释放锁，并进入阻塞状态，等待消息信号
			conder.Wait()
		}

		fmt.Println("满足条件，开始执行")
		conder.L.Unlock()
	}

}
