package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	//swagger:ignore
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func main() {
	var router *gin.Engine = gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Recipe",
		})
	})

	router.Run(":5000")
}
