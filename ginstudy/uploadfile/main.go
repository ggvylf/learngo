package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//上传单个文件
	r.POST("/upload", func(c *gin.Context) {
		//从form表单中获取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		//保存文件，上传到项目根目录
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("file %s is upload,", file.Filename))
	})

	//上传多个文件

	r.Run(":8888")
}
