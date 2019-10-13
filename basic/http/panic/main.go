package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
					<input type="text" name="form1"/>
					<input type="text" name="form2"/>
					<input type="submit" Value="Submit"/>
			</form></body></html>`

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world/n")
	//人为painc
	panic("this is panicsanderrors test")
}

func Form(w http.ResponseWriter, r *http.Request) {
	//添加header信息
	w.Header().Set("Content-Type", "text/html")

	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.Form["form1"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("form2"))
	}
}

func logPanic(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v]caught panicsanderrors:%v", r.RemoteAddr, x)
			}
		}()
		handle(w, r)
	}
}

func main() {

	http.HandleFunc("/hello", logPanic(Hello))
	http.HandleFunc("/form", logPanic(Form))
	url := "127.0.0.1:8888"
	err := http.ListenAndServe(url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
