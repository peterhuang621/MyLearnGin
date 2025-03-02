package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/xxx", "./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		}})
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的播客</a>",
		})
	})
	r.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users/index -- peterhuang",
		})
	})

	// demo - boxer website
	r.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	r.Run(":9090")
}
