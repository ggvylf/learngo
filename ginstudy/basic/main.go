package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//gin.Context封装了request和response
func hello(c *gin.Context) {
	//返回一个json格式
	c.JSON(200, gin.H{
		"message": "hello world!",
	})

	//返回string类型
	c.String(http.StatusOK, "hello world")
}

func main() {

	//创建一个默认路由，同时当中包含Logger和Recovery2个中间件
	r := gin.Default()

	//创建一个默认的路由，不包含中间件
	// r := gin.New()

	//绑定路由规则，方位uri要执行的函数
	r.GET("/hello", hello)

	//启动服务，指定监听端口
	r.Run(":8888")

}
