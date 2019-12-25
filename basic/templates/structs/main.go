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

func sayStruct(w http.ResponseWriter, r *http.Request) {
	person := &Person{"tom", 18}
	tmpl, err := template.ParseFiles("./struct.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, person)

}

func main() {

	http.HandleFunc("/", sayStruct)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
