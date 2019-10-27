package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func CHandlerError(err error, when string) {

	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	udpConn, err := net.Dial("udp", "127.0.0.1:8888")
	CHandlerError(err, "dial")
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)

	buf := make([]byte, 1024)
	for {

		lineBytes, _, _ := reader.ReadLine()

		udpConn.Write(lineBytes)
		n, _ := udpConn.Read(buf)
		fmt.Println("server msg=", string(buf[:n]))

	}

	fmt.Println("client close")

}
