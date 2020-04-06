package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法1 map
		// data := map[string]interface{}{
		// 	"name":    "tom",
		// 	"message": "hello world",
		// 	"age":     18,
		// }

		//方法2 gin.H{} 本质上也是个map
		// data := gin.H{"name": "tom", "message": "hello world", "age": 18}

		//方法3 struct
		type msg struct {
			Name    string `json:"user"`
			Message string `json:"msg"`
			Age     int    `json:"age"`
		}

		data := msg{
			Name:    "tom",
			Message: "hello world",
			Age:     18,
		}

		c.JSON(http.StatusOK, data)
	})

	r.Run(":8888")

}
