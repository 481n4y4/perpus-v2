package main

import (
	"backend/controllers"
	// "backend/middleware"
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
	
	// Auth
	authApi:= router.Group("/api")
	{
		authApi.POST("/auth/register", controllers.Register)
		authApi.POST("/auth/login", controllers.Login)
		authApi.POST("/auth/logout", controllers.Logout)
	}
	// router.POST("/api/auth/register", controllers.Register)
	// router.POST("/api/auth/login", controllers.Login)
	
	
	// CRUD Books
	bookApi := router.Group("/api")
	// bookApi.Use(middleware.RequireAuth())
	{
		bookApi.GET("/books", controllers.FindBooks)
		bookApi.POST("/books", controllers.StoreBook)
		bookApi.GET("/books/:id", controllers.FindBookById)
		bookApi.PUT("/books/:id", controllers.UpdateBook)
		bookApi.DELETE("/books/:id", controllers.DeleteBook)
	}
	// router.GET("/api/books", controllers.FindBooks)
	// router.POST("/api/books", controllers.StoreBook)
	// router.GET("/api/books/:id", controllers.FindBookById)
	// router.PUT("/api/books/:id", controllers.UpdateBook)
	// router.DELETE("/api/books/:id", controllers.DeleteBook)

	router.Run(":8080")
}
