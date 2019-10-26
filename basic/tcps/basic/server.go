package main

import (
	"fmt"
	"net"
	"os"
)

func SHandlerError(err error, when string) {

	if err != nil {
		fmt.Println(when, err)

		os.Exit(1)
	}
}

func ioWithConn(conn net.Conn) {
	buf := make([]byte, 1024)

	for {

		n, err := conn.Read(buf)
		SHandlerError(err, "read")
		clientMsg := string(buf[:n])
		fmt.Println("client msg=", clientMsg)

		if clientMsg == "im off" {
			conn.Write([]byte("bye"))
			break
		}

		conn.Write([]byte("received=" + clientMsg))

	}
	fmt.Println("conn disconnected!")
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	SHandlerError(err, "linstening")
	fmt.Println("listen start")

	for {
		conn, err := listener.Accept()
		SHandlerError(err, "listen Accept")

		go ioWithConn(conn)
	}

	fmt.Println("server stop")
}
