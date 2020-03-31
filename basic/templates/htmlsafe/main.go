package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func xss(w http.ResponseWriter, r *http.Request) {
	//定义自定义函数safe
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")

	if err != nil {
		fmt.Println("template parse failed,err=", err)
		return
	}

	//渲染模板
	str1 := "<script>alert(123)</script>"
	str2 := "<a href=‘http://wwwbaidu.com’>baidu</a>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Println("template exec failed err=", err)
	}

}

func main() {
	//创建http服务
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("http server start failed,err=", err)
		return
	}

}
