package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLFiles("./xss.tmpl")

	r.GET("/safepage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "xss.tmpl", gin.H{
			"str1": "<script>alert(123)</script>",
			"str2": "<a href=‘http://wwwbaidu.com’>baidu</a>",
		})
	})

	r.Run((":8888"))

}
