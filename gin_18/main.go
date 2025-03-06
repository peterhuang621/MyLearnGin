package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler")
	name, ok := c.Get("name")
	if !ok {
		name = "anonymous"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	//func(c.Copy())....
	c.Next()
	// c.Abort()
	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "peter")
	c.Next()
	// c.Abort()
	fmt.Println("m2 out...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if doCheck {
			ctx.Next()
		} else {
			ctx.Abort()
		}
	}
}

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(m1, m2)
	r.GET("/index", indexHandler)
	r.GET("/shop", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})

	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}

	xx2Group := r.Group("/xx2", authMiddleware(true))
	{
		xx2Group.GET("index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
		})
	}
	r.Run(":9090")
}
