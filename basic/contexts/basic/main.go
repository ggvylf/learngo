package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 方法1 全局变量
// var wg sync.WaitGroup
// var notify bool

// func f() {
// 	defer wg.Done()
// 	for {
// 		fmt.Println("hehe")
// 		time.Sleep(500 * time.Millisecond)
// 		if notify {
// 			break
// 		}
// 	}

// }

// func main() {
// 	wg.Add(1)
// 	go f()
// 	time.Sleep(3 * time.Second)
// 	notify = true
// 	wg.Wait()
// }

// 方法2 使用channel
// var wg sync.WaitGroup
// var exitChan = make(chan bool)

// func f() {
// 	defer wg.Done()
// 	for {
// 		fmt.Println("hehe")
// 		time.Sleep(500 * time.Millisecond)
// 		select {
// 		case <-exitChan:
// 			break
// 		default:
// 		}
// 	}

// }

// func main() {
// 	wg.Add(1)
// 	go f()
// 	time.Sleep(3 * time.Second)
// 	exitChan <- true
// 	wg.Wait()
// }

// 方法3 使用LOOP跳出循环
// LOOP:
// func f() {
// 	defer wg.Done()
// 	for {
// 		fmt.Println("hehe")
// 		time.Sleep(500 * time.Millisecond)
// 		select {
// 		case <-exitChan:
// 			break LOOP
// 		default:
// 		}
// 	}

//方法4  使用context
var wg sync.WaitGroup

func f1(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
	for {
		fmt.Println("f1")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}

}

func f2(ctx context.Context) {
	defer wg.Done()
	for {
		fmt.Println("f2")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}

}

func main() {
	// 创建一个context、
	//context.WithCancel()返回一个带有新Done chan的父节点的副本。当调用cancel()或父上下文的Done chan时，将关闭返回上下文的Done chan。
	//取消该context将释放与其相关的资源
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go f1(ctx)
	time.Sleep(3 * time.Second)
	//执行cancel方法，往ctx.Down()这个chan中传递信号，通知gorutine结束
	cancel()
	wg.Wait()
}
