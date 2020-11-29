package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//全局中间件
//所有的请求都经过
func GlobalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		fmt.Println("globalmiddleware is starting")

		//设置变量到Context中，可以使用c.Get()获取
		c.Set("request", "globalmiddlewarevalue")

		//执行其他函数
		fmt.Println("globalmiddleware starting running other func")
		c.Next()

		//执行其他函数后，中间件继续执行后续的操作
		status := c.Writer.Status()
		fmt.Println("globalmiddleware complated,status=", status)

		t2 := time.Since(t1)
		fmt.Println("globalmiddleware running time=", t2)
	}
}

func LocalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("localmiddleware is running")
		fmt.Println("localmiddleware starting running other func")
		c.Next()

		status := c.Writer.Status()
		fmt.Println("localmiddleware complated ,http code=", status)

	}
}

func main() {
	r := gin.Default()

	//注册中间件
	r.Use(GlobalMiddleWare())

	{
		//调用全局中间件
		r.GET("/globalmiddleware", func(c *gin.Context) {

			//从中间件中获取key对应的value
			value, _ := c.Get("request")
			fmt.Println("middleware key request value=", value)

			//页面显示
			c.JSON(200, gin.H{"request": value})
		})

		//调用局部中间件，调用handler的时候传入即可
		r.GET("/localmiddleware", LocalMiddleWare(), func(c *gin.Context) {

			//页面显示
			c.JSON(200, gin.H{"localmiddleware": "ok"})
		})
	}

	r.Run(":8888")
}
