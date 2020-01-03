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
	person1 := &Person{"tom", 18}
	person2 := map[string]interface{}{
		"name": "jerry",
		"age":  13,
	}

	hobby := []string{"aaaa", "bbbb", "cccc"}

	users := map[string]interface{}{
		"person1": person1,
		"person2": person2,
		"hobby":   hobby,
	}
	tmpl, err := template.ParseFiles("./struct.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, users)

}

func main() {

	http.HandleFunc("/", sayStruct)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
