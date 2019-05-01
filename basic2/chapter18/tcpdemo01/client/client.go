package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接失败，错误为=", err)
		return
	}
	fmt.Println("连接服务器成功，conn=", conn)

	//从终端中读取数据,Stdin是一个文件描述符
	reader := bufio.NewReader(os.Stdin)

	for {
		//从终端读取一行内容，结束符为换行
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取错误，错误为", err)
		}

		//去掉line结尾的换行符
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("正常退出")
			return
		}

		//把数据发送给服务器，并返回数据长度
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("写入错误，错误为", err)
		}
		fmt.Printf("客户发送了%dbyte的数\n", n)
	}
}
