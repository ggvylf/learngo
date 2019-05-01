package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	//json
	// conffile := "./test.json"
	// conftype := "json"

	conffile := "./test.ini"
	conftype := "ini"

	conf, err := config.NewConfig(conftype, conffile)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	addr := conf.String("server::listen_ip")

	fmt.Println("ip:", addr)

	port, err := conf.Int("server::listen_port")
	if err != nil {
		fmt.Println("read server:port failed, err:", err)
		return
	}

	fmt.Println("port:", port)

	log_level := conf.String("logs::log_level")
	if len(log_level) == 0 {
		log_level = "debug"
	}

	fmt.Println("log_level:", log_level)

	log_path := conf.String("collect::log_path")
	fmt.Println("log_path:", log_path)
}
