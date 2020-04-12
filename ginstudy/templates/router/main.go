package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//没有匹配的到路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not route not found",
		})
	})

	//指定某个url的全部方法
	r.Any("/test", func(c *gin.Context) {

		//用switch case来处理
		switch c.Request.Method {
		//这里直接使用http定影好的const
		case http.MethodGet:
			{
				c.JSON(http.StatusOK, gin.H{
					"message": "method get",
				})
			}
		case http.MethodPost:
			{
				c.JSON(http.StatusOK, gin.H{
					"message": "method post",
				})
			}
		case http.MethodPut:
			{
				c.JSON(http.StatusOK, gin.H{
					"message": "method put",
				})
			}
		case http.MethodDelete:
			{
				c.JSON(http.StatusOK, gin.H{
					"message": "method delete",
				})
			}
		}
	})

	//group route 支持多级嵌套
	video := r.Group("/video")
	{
		video.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/login",
			})
		})

		video.GET("/order", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/order",
			})
		})
	}

	book := r.Group("/book")
	{
		book.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/book/login",
			})
		})

		book.GET("/order", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/book/order",
			})
		})
	}

	r.Run(":8888")
}
