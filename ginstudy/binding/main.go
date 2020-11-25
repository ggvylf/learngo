package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required" `
	Passwd   string `form:"passwd" json:"passwd" uri:"passwd" xml:"passwd" binding:"required" `
}

func main() {
	r := gin.Default()

	//json绑定
	//测试：curl -v -X POST http://127.0.0.1:8888/loginJSON -H 'content-type: application/json'  -d '{ "username": "tom",passwd":"123" }'
	r.POST("/loginJSON", func(c *gin.Context) {
		//声明接收传参的结构体
		var json Login
		//把request.body中的数据按照json格式解析
		if err := c.ShouldBindJSON(&json); err != nil {
			//解析失败报错
			//gin.H{}封装了生成json的数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//判断接收到的参数
		if json.Username != "tom" || json.Passwd != "123" {
			c.JSON(http.StatusBadRequest, gin.H{"http code": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"http code": "200"})
		c.String(http.StatusOK, fmt.Sprintf("username:%s--passwd:%s\n", json.Username, json.Passwd))

	})

	//form表单
	r.POST("loginForm", func(c *gin.Context) {
		var form Login
		//Bind()默认解析为form格式。根据Context-Type来自动判断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//判断接收到的参数
		if form.Username != "tom" || form.Passwd != "123" {
			c.JSON(http.StatusBadRequest, gin.H{"http code": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"http code": "200"})
		c.String(http.StatusOK, fmt.Sprintf("username:%s--passwd:%s\n", form.Username, form.Passwd))

	})

	//uri参数
	r.GET("loginURI/:username/:passwd", func(c *gin.Context) {
		var uri Login

		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//判断接收到的参数
		if uri.Username != "tom" || uri.Passwd != "123" {
			c.JSON(http.StatusBadRequest, gin.H{"http code": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"http code": "200"})
		c.String(http.StatusOK, fmt.Sprintf("username:%s--passwd:%s\n", uri.Username, uri.Passwd))
	})

	r.Run(":8888")
}
