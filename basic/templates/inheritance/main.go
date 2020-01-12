package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./index.tmpl")
	if err != nil {
		fmt.Println("parse template err,err=", err)
		return
	}
	message := "this is index template"
	t.ExecuteTemplate(w, "index.tmpl", message)
}

func hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./hello.tmpl")
	if err != nil {
		fmt.Println("parse template err,err=", err)
		return
	}
	message := "this is hello template"
	t.ExecuteTemplate(w, "hello.tmpl", message)
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("server listen err,err=", err)
	}
}
