package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {

		//创建一个空slice
		buf := make([]byte, 1024)

		//从conn中读取数据，如果客户端没有write，就阻塞
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取错误，err=", err)
			return
		}
		fmt.Printf("接收%v的%dbyte的数据=%v\n", conn.RemoteAddr().String(), n, string(buf[:n]))
	}

}

func main() {
	fmt.Println("服务器开始监听")

	//创建一个tcp协议的监听器实例，ip和端口是本机的8888
	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println(err)
		return
	}
	//延时关闭连接
	defer listen.Close()

	for {
		fmt.Println("等待客户端连接。。。")

		//接收请求内容
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)

		} else {
			//获取请求的ip地址
			fmt.Printf("Accept() success conn，客户端地址=%v\n", conn.RemoteAddr())

		}
		go process(conn)
	}

}
