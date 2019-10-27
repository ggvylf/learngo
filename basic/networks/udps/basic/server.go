package main

import (
	"fmt"
	"net"
	"os"
)

func SHandlerError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)

		os.Exit(1)
	}
}

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	SHandlerError(err, "resolv client ip")

	udpConn, err := net.ListenUDP("udp", udpAddr)
	SHandlerError(err, "get conn")
	defer udpConn.Close()

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := udpConn.ReadFromUDP(buf)
		SHandlerError(err, "read msg")

		fmt.Printf("remote addr=%v,msg=%s\n", remoteAddr, string(buf[:n]))
		udpConn.WriteToUDP([]byte("rev="+string(buf[:n])), remoteAddr)
	}
}
