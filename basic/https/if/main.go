package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
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

	var arr []Person

	p1 := Person{Name: "Mary", Age: 13}
	p2 := Person{Name: "tom", Age: 22}
	p3 := Person{Name: "jerry", Age: 23}
	p4 := Person{Name: "mark", Age: 44}

	arr = append(arr, p1)
	arr = append(arr, p2)
	arr = append(arr, p3)
	arr = append(arr, p4)

	myTemplate.Execute(w, arr)

}

func main() {
	//初始化模板
	//filename := "./index.html"
	//filename := "./with.html"
	filename := "./range.html"
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
