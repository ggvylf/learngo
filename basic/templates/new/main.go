package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	mylen := func(a string) (int, error) {
		s
		return len(a), nil
	}

	t, err := template.New("mytmpl.tmpls").ParseFiles("./mytmpl.tmpl")

	if err != nil {
		fmt.Println("parse new template failed!,err=", err)
		return
	}
	name := "tom"
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
