package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//auth中间件
func AuthMiddleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("mycookie")

		//验证通过
		if err == nil && cookie == "tom123" {
			//继续执行其他函数
			c.Next()
			return
		}
		//验证不通过，不再调用后续程序
		c.JSON(401, gin.H{"error": "no auth"})
		c.Abort()
		return
	}

}

func main() {
	r := gin.Default()

	//调用auth中间件
	r.GET("/home", AuthMiddleHandler(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "this is my home!"})

	})

	r.GET("/login", func(c *gin.Context) {
		username := c.Query("username")
		passwd := c.Query("passwd")

		if username != "tom" || passwd != "123" {
			c.JSON(401, gin.H{"error": "username or passwd is wrong!"})
			return
		}
		fmt.Printf("username=%s,passwd=%s\n", username, passwd)
		//用户名密码合规，创建cookies
		c.SetCookie("mycookie", fmt.Sprintf("%s%s", username, passwd), 60, "/", "172.18.203.74", false, true)

	})

	r.Run(":8888")

}
