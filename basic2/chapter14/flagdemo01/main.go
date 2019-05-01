package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "p", "", "密码，默认为空")
	flag.StringVar(&host, "h", "", "主机，默认为空")
	flag.IntVar(&port, "port", 0, "端口，默认为0")

	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v\n", user, pwd, host, port)

}
