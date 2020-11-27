package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 异步处理
	r.GET("/long_async", func(c *gin.Context) {

		//创建一个上下文的副本
		copyContext := c.Copy()

		go func() {
			time.Sleep(3 * time.Second)
			log.Println("async running" + copyContext.Request.URL.Path)
		}()

	})

	//同步处理
	r.GET("long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("sync running" + c.Request.URL.Path)
	})

	r.Run(":8888")
}
