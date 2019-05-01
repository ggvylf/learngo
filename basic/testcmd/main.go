package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup

func OsCmd(host string, cmd string) {
	defer wg.Done()
	run := exec.Command("/usr/bin/ssh", host, cmd)
	//run := exec.Command("/usr/bin/ssh", "root@192.168.56.21", "/usr/bin/uptime")

	//获取输出
	var out bytes.Buffer
	run.Stdout = &out

	//执行命令
	err := run.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(run.Stdout)
}

func main() {

	if len(os.Args[1:]) < 2 {
		fmt.Println("Usage testcmd cmd host1 host2 ...")
		os.Exit(1)
	}
	hosts := os.Args[2:]
	cmd := os.Args[1]
	wg.Add(len(hosts))
	for _, host := range hosts {
		go OsCmd(host, cmd)

	}
	wg.Wait()
}
