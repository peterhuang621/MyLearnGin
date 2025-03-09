package routers

import (
	"gin_demo/demo_25/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.CreateTodo)

		v1Group.GET("/todo", controller.GetTodoList)

		v1Group.GET("/todo/query", controller.GetTodo_Query)

		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
