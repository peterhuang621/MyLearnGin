package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", sayHello)

	r.POST("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "POST"})
	})

	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "PUT"})
	})

	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
	})

	r.Run(":9090")
}
