package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//打开模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("template parse failed,err=", err)
		return
	}

	//渲染模板
	str := "World"
	err = t.Execute(w, str)
	if err != nil {
		fmt.Println("template exec failed err=", err)
	}

}

func main() {
	//创建http服务
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("http server start failed,err=", err)
		return
	}

}
