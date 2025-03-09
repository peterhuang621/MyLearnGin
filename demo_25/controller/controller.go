package controller

import (
	"gin_demo/demo_25/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"msg": "ok"})
}

func CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	ctx.BindJSON(&todo)
	err := models.CreateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(ctx *gin.Context) {
	todoList, err := models.GetAllTodo()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "invalid id!"})
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	ctx.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "invalid id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

func GetTodo_Query(ctx *gin.Context) {
	todoList, err := models.GetATodo(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}
