package main

import (
	"github.com/gin-gonic/gin"
)

func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(200, "book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.LoadHTMLGlob("./book/templates/*")

	r.GET("/book/list", bookListHandler)

	r.Run(":8888")
}
