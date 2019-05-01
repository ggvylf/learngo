package main

import (
	"fmt"
	"net/http"
)

//handler对应的方法
func Index(w http.ResponseWriter, r *http.Request) {
	//服务端打印内容
	fmt.Println("index")

	//客户端返回内容
	fmt.Fprint(w, "index\n")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	fmt.Fprint(w, "hello\n")
}

func main() {
	//定义handler和对应的方法
	http.HandleFunc("/", Index)
	http.HandleFunc("/hello", Hello)

	//创建http监听，本质上是封装了tcp
	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		fmt.Println("listen failed")
	}

}
