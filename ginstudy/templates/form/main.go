package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.LoadHTMLFiles("./login.tmpl", "./result.tmpl")

	//get请求是用来显示页面的
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", nil)
	})

	//post请求是用来处理提交的数据
	r.POST("/login", func(c *gin.Context) {

		//方法1 获取提交的数据
		// username := c.PostForm("username")
		// password := c.PostForm("password")

		//方法2 制定默认值
		// username := c.DefaultPostForm("username", "tom")
		// password := c.DefaultPostForm("password", "123")

		// 方法3 判断ok
		username,ok := c.PostForm("username")
		if !ok {
			username="tom"
		}

		password,ok := c.PostForm("password")
		if ！ok {
			password="123"
		}		



		c.HTML(http.StatusOK, "result.tmpl", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":8888")
}
