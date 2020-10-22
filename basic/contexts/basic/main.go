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

func f(ctx context.Context) {
	defer wg.Done()
	for {
		fmt.Println("hehe")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}

}

func main() {
	// 创建一个context
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go f(ctx)
	time.Sleep(3 * time.Second)
	//执行cancel方法，往ctx.Down()这个chan中传递信号
	cancel()
	wg.Wait()
}
