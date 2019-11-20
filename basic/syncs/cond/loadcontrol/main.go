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

func NewSever(maxconn int) *Server {
	server := new(Server)
	server.maxconn = maxconn
	server.cond = sync.NewCond(&sync.Mutex{})
	server.chAlarm = make(chan bool, 1)
	return server
}

func (server *Server) watchAlarm() {
	go func() {
		for {
			<-server.chAlarm

			server.cond.L.Lock()
			// <-time.After(3 * time.Second)
			server.conn--
			fmt.Println("load --,current conn=", server.conn)
			fmt.Println("load is ok")
			server.cond.Broadcast()
			server.cond.L.Unlock()
		}
	}()

}

func (server *Server) Run() {
	for {

		server.cond.L.Lock()
		for server.conn == server.maxconn {
			server.chAlarm <- true
			fmt.Println("much too load,sent alarm")
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

	server.watchAlarm()
	server.Run()

}
