package main

import "github.com/gin-gonic/gin"

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world!",
	})
}

func main() {

	//创建一个默认路由
	r := gin.Default()

	//制定路由
	r.GET("/hello", hello)

	//启动服务
	r.Run(":8888")

}
