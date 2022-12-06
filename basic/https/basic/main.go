package main

import (
	"fmt"
	"net/http"
	"time"
)

// http包route的匹配规则
// 1. 结尾是/xxx，表示一个子树，后边可以跟其他路径，末尾不是/ ,表示一个固定的路径
// 2. 以/结尾的URL可以匹配它的任何子路径， 例如/images 会匹配到/images/a.jpg
// 3. 最长匹配原则，如果有多个匹配，会使用路径最长的来处理
// 4. 找不到任何匹配项 返回404

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world ! \n")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timestr := fmt.Sprintf("\"time\": \"%s\" \n", t)

		w.Write([]byte(timestr))
	})

	// 访问 /time2 /time2/1 /time2/xxx都能返回
	http.HandleFunc("/time2/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timestr := fmt.Sprintf("\"time\": \"%s\" \n", t)

		w.Write([]byte(timestr))
	})

	// 启动http服务，指定监听的端口
	http.ListenAndServe(":8888", nil)
}
