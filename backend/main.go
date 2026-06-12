package main

import (
	"backend/controllers"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/api/books", controllers.FindBooks)
	router.POST("/api/books", controllers.StoreBook)
	router.GET("/api/books/:id", controllers.FindBookById)
	router.PUT("/api/books/:id", controllers.UpdateBook)
	router.DELETE("/api/books/:id", controllers.DeleteBook)

	router.Run(":3000")
}
