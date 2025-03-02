package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type msg struct {
	Name    string `json:"json_name"`
	Message string
	Age     int
}

func main() {
	r := gin.Default()
	r.GET("/json", func(ctx *gin.Context) {
		// data := map[string]interface{}{
		// 	"name":    "peterhuang",
		// 	"message": "hello world",
		// 	"age":     18,
		// }
		data := gin.H{
			"name":    "peterhuang",
			"message": "hello world",
			"age":     18,
		}

		ctx.JSON(http.StatusOK, data)
	})

	r.GET("/another_json", func(ctx *gin.Context) {
		data := msg{"小王子", "hello golang", 20}
		ctx.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
