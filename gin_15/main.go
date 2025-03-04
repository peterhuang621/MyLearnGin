package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(ctx *gin.Context) {
		f, err := ctx.FormFile("f1")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			dst := fmt.Sprintf("./%s", f.Filename)
			ctx.SaveUploadedFile(f, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("%s is saved!", dst),
			})
		}
	})
	r.Run(":8080")
}
