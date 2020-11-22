package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHello(c *gin.Context) {

	c.String(http.StatusOK, "get hello world")
}

func postHello(c *gin.Context) {

	c.String(http.StatusOK, "post hello world")
}

func putHello(c *gin.Context) {

	c.String(http.StatusOK, "put hello world")
}

func deleteHello(c *gin.Context) {

	c.String(http.StatusOK, "delete hello world")
}

func main() {
	r := gin.Default()

	r.GET("/hello", getHello)
	r.POST("/hello", postHello)
	r.PUT("/hello", putHello)
	r.DELETE("/hello", deleteHello)

	r.Run(":8888")

}
