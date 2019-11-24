package main

import (
	"fmt"
	"sync"
	"time"
)

type Server struct {
	conn    int
	maxconn int
	chAlarm chan bool
	cond    *sync.Cond
}

//工厂模式
func NewSever(maxconn int) *Server {
	server := new(Server)
	server.maxconn = maxconn
	//创建条件，在互斥锁保护下
	server.cond = sync.NewCond(&sync.Mutex{})
	//使用带缓存的chan，防止空chan读取panic
	server.chAlarm = make(chan bool, 1)
	return server
}

func (server *Server) watchAlarm() {
	go func() {
		for {
			//chan中只要有数据，表示已经达到阈值
			<-server.chAlarm

			//操作数据之前要加锁
			server.cond.L.Lock()
			// <-time.After(3 * time.Second)
			server.conn--
			fmt.Println("load --,current conn=", server.conn)
			fmt.Println("load is ok")
			//发送信号
			server.cond.Broadcast()
			//解锁
			server.cond.L.Unlock()
		}
	}()

}

func (server *Server) Run() {
	for {

		server.cond.L.Lock()
		for server.conn == server.maxconn {
			//到达阈值，写入chan
			server.chAlarm <- true
			fmt.Println("much too load,sent alarm")
			//进入等待阻塞状态，等待goroutine的信号
			server.cond.Wait()

		}
		time.Sleep(1 * time.Second)
		server.conn++
		fmt.Println("load ++,current conn=", server.conn)
		server.cond.L.Unlock()
	}

}

func main() {
	server := NewSever(3)

	//一定要让读取chan的函数再在前边执行
	server.watchAlarm()
	server.Run()

}
