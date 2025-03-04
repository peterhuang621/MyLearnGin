package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(ctx *gin.Context) {
		// name := ctx.Query("query")
		// name := ctx.DefaultQuery("query", "somebody")
		name, ok := ctx.GetQuery("query")
		age := ctx.Query("age")
		if !ok {
			name = "somebody"
		}
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}
