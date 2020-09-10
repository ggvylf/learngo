package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//跟server建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dail failed,err=", err)
		return
	}

	//发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("input someting,exit to quit: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))

	}

	conn.Close()
}
