package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载模板文件
	// r.LoadHTMLFiles("./templates/index.tmpl")

	//加载模板文件，通过通配符匹配的方式
	r.LoadHTMLGlob("./templates/*")

	//渲染模板文件
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			//替换index.tmpl中的{{ .title }}
			"title": "test gin",
		})
	})

	r.Run(":8888")

}
