package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//HTTP重定向
	r.GET("/index", func(c *gin.Context) {
		//301重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	//路由重定向
	r.GET("/a", func(c *gin.Context) {
		//改写请求的uri
		c.Request.URL.Path = "/b"

		//路由继续处理后续请求
		r.HandleContext(c)
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is b page",
		})
	})

	r.Run(":8888")
}
