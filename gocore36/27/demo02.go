package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

	sign := make(chan struct{}, 1)
	max := 5

	//发信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 {
				log.Printf("sender [%d]: sender is waiting", i)
				sendCond.Wait()
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			//给接收端发消息，可以取信了。
			recvCond.Signal()
		}
	}(max)

	//接收端
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			if mailbox == 0 {
				log.Printf("receiver [%d]: recevier is waiting", j)
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been received.", j)

			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
