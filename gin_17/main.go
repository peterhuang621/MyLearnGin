package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"method": "Get",
		})
	})
	r.POST("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"method": "Post",
		})
	})
	r.PUT("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "Put",
		})
	})
	r.DELETE("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "Delete",
		})
	})
	// r.Any("/user", func(ctx *gin.Context) {
	// 	switch ctx.Request.Method {
	// 	case http.MethodGet:
	// 		ctx.JSON(http.StatusOK, gin.H{"method": "GET"})
	// 	case http.MethodPost:
	// 		ctx.JSON(http.StatusOK, gin.H{"method": "POST"})
	// 	}
	// })
	videogroup := r.Group("/video")
	// r.GET("/video/index", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	// })
	// r.GET("/video/cc", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"msg": "/video/cc"})
	// })
	// r.GET("/video/bb", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"msg": "/video/bb"})
	// })
	{
		videogroup.GET("index", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "/video/index"}) })
		videogroup.GET("cc", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "/video/cc"}) })
		videogroup.GET("bb", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "/video/bb"}) })
	}

	r.GET("/shop/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	})
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})
	r.Run(":8080")
}
