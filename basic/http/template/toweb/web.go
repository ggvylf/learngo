package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  string
}

type Result struct {
	output string
}

//实现 io.Writer方法
func (r *Result) Write(b []byte) (n int, err error) {
	fmt.Println("call by template")
	r.output += string(b)
	return len(b), nil
}

var myTemplate *template.Template

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return

}

func UserInfo(w http.ResponseWriter, r *http.Request) {

	p := Person{Name: "Mary", Age: "31"}

	// file, err := os.OpenFile("./test.out", os.O_CREATE|os.O_WRONLY, 0755)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	result := &Result{}

	//通过传递不同的io.Writer，是多态的实现
	myTemplate.Execute(result, p)
	fmt.Println(result.output)

}

func main() {
	//初始化模板
	filename := "./index.html"
	initTemplate(filename)

	//handler
	http.HandleFunc("/userinfo", UserInfo)

	//web server
	url := "127.0.0.1:8888"
	err := http.ListenAndServe(url, nil)
	if err != nil {
		fmt.Println(err)
	}

}
