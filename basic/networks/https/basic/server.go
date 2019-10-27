package main

import "net/http"

func myTest(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("this is a test routing\n"))
}

func myHello(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Add("Test-Header", "hehe")
}

func main() {

	http.HandleFunc("/test", myTest)
	http.HandleFunc("/hello", myHello)
	http.ListenAndServe("127.0.0.1:8888", nil)
}
