package main

import "github.com/gin-gonic/gin"

func main() {
	var router *gin.Engine = gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin Apps",
		})
	})

	router.Run(":5000")
}
