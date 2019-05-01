package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:8888"

	//使用get方法获取请求结果
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("get url err")
		return
	}

	//读取请求结果的body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read body err")
		return
	}

	fmt.Println(string(data))

}
