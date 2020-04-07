package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//指定内存限制，默认是32MiB,这里调整为８MiB
	r.MaxMultipartMemory = 8 << 20

	r.LoadHTMLFiles("./index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// //上传单个文件
	// r.POST("/upload", func(c *gin.Context) {
	// 	file, err := c.FormFile("f1")
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": err.Error(),
	// 		})

	// 	}
	// 	//指定保存的路径
	// 	//方法１
	// 	// dst:=fmt.Sprintf("./",file.Filename)
	// 	//方法２
	// 	dst := path.Join("./", file.Filename)

	// 	c.SaveUploadedFile(file, dst)
	// })

	// 上传多个文件
	r.POST("/upload", func(c *gin.Context) {

		form, _ := c.MultipartForm()
		files := form.File["f1"]

		for _, file := range files {
			dst := path.Join("./", file.Filename)
			c.SaveUploadedFile(file, dst)
		}

	})

	r.Run(":8888")
}
