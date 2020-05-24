package main

import (
	"github.com/dkzhang/RmsGo/webapi/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerGinToFile())
	r.GET("/user/*action", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":   "hello",
			"action": c.Param("action"),
		})
	})
	r.Run()
}
