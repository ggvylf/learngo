package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func mydelies(w http.ResponseWriter, h *http.Request) {
	//定义新的标识符
	t, err := template.New("base.tmpl").Delims("{[", "]}").ParseFiles("./base.tmpl")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	msg := "hello"
	t.Execute(w, msg)

}

func main() {
	http.HandleFunc("/", mydelies)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
}
