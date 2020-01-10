package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// //自定义一个函数
	// mylen := func(a string) (int, error) {
	// 	return len(a), nil
	// }

	// //初始化模板
	// t := template.New("./mytmpl.html")
	// t.Funcs(template.FuncMap{
	// 	"mylen": mylen,
	// })

	//解析模板
	t, err := template.ParseFiles("./mytmpl.tmpl", "./my1.tmpl")
	if err != nil {
		fmt.Println("parse template file failed,err=", err)
		return
	}

	name := "tom"
	//渲染模板
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("server listen failed,err=", err)
		return
	}
}
