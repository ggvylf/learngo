package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func process() {
	//生成context实例和超时定时器
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//取消定时器
	defer cancel()

	//自己实现http.Client
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	//初始化管道 存放http resp的结果
	resultchan := make(chan Result, 1)

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		//执行req
		resp, err := client.Do(req)

		//实例化result
		pack := Result{r: resp, err: err}
		//存入管道
		resultchan <- pack

	}()
	select {
	//超时管道中有数据
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-resultchan
		fmt.Println("timeout!", res)
		return

		//管道中有数据
	case res := <-resultchan:

		out, _ := ioutil.ReadAll(res.r.Body)
		defer res.r.Body.Close()

		fmt.Println("server res=", out)
	}
	return

}

func main() {
	process()
}
