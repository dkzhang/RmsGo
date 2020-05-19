package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/*action", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":   "hello",
			"action": c.Param("action"),
		})
	})
	r.Run()
}
