package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func myGet() {
	domain := "https://www.baidu.com"
	query := "/s?wd=golang"
	url := domain + query

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

func myPost() {
	domain := "https://httpbin.org"
	query := "/post?name=tom&age=18"
	url := domain + query

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("aa=1&bbb=2"))
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

func main() {
	myGet()
	myPost()
}
