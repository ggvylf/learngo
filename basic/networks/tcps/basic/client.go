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
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	CHandlerError(err, "dial")

	reader := bufio.NewReader(os.Stdin)

	buf := make([]byte, 1024)
	for {


		lineBytes, _, _ := reader.ReadLine()
		conn.Write(lineBytes)

		n, _ := conn.Read(buf)
		serverMsg := string(buf[:n])
		fmt.Println("server msg=", serverMsg)

		if serverMsg == "bye" {
			break
		}
	}

	fmt.Println("client close")

}
