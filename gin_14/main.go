package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/user", func(ctx *gin.Context) {
		// username := ctx.Query("username")
		// password := ctx.Query("password")
		// u := UserInfo{
		// 	username: username,
		// 	password: password,
		// }
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.POST("/form", func(ctx *gin.Context) {
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.POST("/json", func(ctx *gin.Context) {
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
		// var u UserInfo
		// err := ctx.ShouldBind(&u)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{
		// 		"error": err.Error(),
		// 	})
		// } else {
		// 	fmt.Printf("%#v\n", u)
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"message": "ok",
		// 	})
		// }
	})
	r.Run(":9090")
}
