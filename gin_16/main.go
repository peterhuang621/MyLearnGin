package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"status": "ok",
		// })
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	r.GET("/a", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/b"
		r.HandleContext(ctx)
	})
	r.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	r.Run(":8080")
}
