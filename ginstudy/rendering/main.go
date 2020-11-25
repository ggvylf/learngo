package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	r := gin.Default()

	//json响应
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})

	//sturct响应
	r.GET("/someSTRUCT", func(c *gin.Context) {
		var msg struct {
			Message string
			Status  int
		}
		msg.Message = "someSTRUCT"
		msg.Status = 200

		c.JSON(200, msg)

	})

	//xml响应
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "someJSON", "status": 200})
	})

	//yaml响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"message": "someYAML", "status": 200})
	})

	//protobuf响应
	r.GET("/somePROTOBUF", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "somePROTOBUF"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)

	})

	r.Run(":8888")
}
