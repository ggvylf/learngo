package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//uri参数
	r.GET("/uriparam/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.String(http.StatusOK, name+age+"\n")
	})

	//星号会把action后边的所有内容当成一个
	r.GET("/urimaram2/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" - "+strings.Split(action, "/")[1]+"\n")
	})

	//api参数
	//注意curl的时候有&符号要加引号
	r.GET("/apiuri", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		c.String(http.StatusOK, name+age+"\n")
	})

	//key不存在就返回默认值
	r.GET("/apiuridefault", func(c *gin.Context) {
		name := c.DefaultQuery("name", "zhangsan")
		age := c.DefaultQuery("age", "18")
		c.String(http.StatusOK, name+age+"\n")
	})

	//form表单
	r.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert")
		//单个值
		username := c.PostForm("username")
		passwd := c.PostForm("passwd")
		//多选框
		hobbys := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf("type1=%s, username=%s,passwd=%s,hobbys=%s\n", type1, username, passwd, hobbys))
	})

	r.Run(":8888")

}
