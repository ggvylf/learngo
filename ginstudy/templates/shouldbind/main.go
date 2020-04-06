package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username" form:"username" `
	Password string `json:"password" form:"password" `
}

func main() {

	r := gin.Default()

	//querystring
	r.GET("/loginquery", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"password": user.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	//from表单
	r.POST("/loginform", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"password": user.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	//json
	r.POST("/loginjson", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"password": user.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.Run(":8888")
}
