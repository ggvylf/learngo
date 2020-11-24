package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8888")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "mylogin")
	c.String(200, fmt.Sprintf("name=%s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "mysubmit")
	c.String(200, fmt.Sprintf("name=%s\n", name))
}
