package main

import (
	"fmt"
	"net"
)

func processConn(conn net.Conn) {
	//跟client通信
	var tmp [128]byte

	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed,err=", err)
		}
		fmt.Println(string(tmp[:n]))
	}
}

func main() {

	//启动服务建立监听
	listener, err := net.Listen("tcp", "127.0.0.01:8888")
	if err != nil {
		fmt.Println("listen failed", err)

	}

	for {
		//等待client的链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			return
		}

		go processConn(conn)

	}

}
