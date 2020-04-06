package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//获取请求参数
		// name := c.Query("name")
		// age := c.Query("age")

		//指定默认值
		// name := c.DefaultQuery("name", "jerry")
		// age := c.DefaultQuery("age", "18")

		//通过ok判断
		name, ok := c.GetQuery("name")
		if !ok {
			name = "jerry"
		}

		age, ok := c.GetQuery("age")
		if !ok {
			age = "18"
		}

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8888")
}
