package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(ctx *gin.Context) {
		// username := ctx.PostForm("username")
		// password := ctx.PostForm("password")
		username := ctx.DefaultPostForm("username", "somebody")
		password := ctx.DefaultPostForm("password", "somepassword")
		// password := ctx.DefaultPostForm("xxx", "somepassword") retrieve somepassword
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":9090")
}
