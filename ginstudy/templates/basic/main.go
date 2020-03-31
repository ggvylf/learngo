package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//解析模板
	r.LoadHTMLFiles("./templates/index.tmpl")

	//渲染模板
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "test gin",
		})
	})

	r.Run(":8888")

}
