package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)



func main() {
	urls := []string{
		"http://www.baidu.com",
		"http://www.163.com",
		"http://hehe.com",
		"http://google.com",
	}

	for _, v := range urls {
		client := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}

		res, err := client.Head(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		code := res.StatusCode
		fmt.Println(v, "code=", code)

	}

}
