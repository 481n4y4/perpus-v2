package main

import (
	"gihub.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world"
		})
	})

	router.Run(":3000")
}