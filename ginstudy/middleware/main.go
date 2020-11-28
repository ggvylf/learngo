package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//全局中间件
//所有的请求都经过
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		fmt.Println("middleware is starting")

		//再中间件中设置key和value
		c.Set("request", "middlewarevalue")

		//执行其他函数
		fmt.Println("starting running other func")
		c.Next()

		//执行其他函数后，中间件继续执行后续的操作
		status := c.Writer.Status()
		fmt.Println("middleware complated,status=", status)

		t2 := time.Since(t1)
		fmt.Println("middleware running time=", t2)
	}
}

func main() {
	r := gin.Default()

	//注册中间件
	r.Use(MiddleWare())
	{
		r.GET("/middleware", func(c *gin.Context) {

			//从中间件中获取key对应的value
			value, _ := c.Get("request")
			fmt.Println("middleware key request value=", value)

			//页面显示
			c.JSON(200, gin.H{"request": value})
		})
	}

	r.Run(":8888")
}
