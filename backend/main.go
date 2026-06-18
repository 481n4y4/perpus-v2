package main

import (
	"backend/controllers"
	"flag"
	"time"

	// "backend/middleware"
	"backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	fresh := flag.Bool("fresh", false, "Hapus dan buat ulang semua tabel")
	flag.Parse()

	models.ConnectDatabase(*fresh)

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},

		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Auth
	authApi := router.Group("/api/auth")
	{
		authApi.POST("/register", controllers.Register)
		authApi.POST("/login", controllers.Login)
		authApi.POST("/logout", controllers.Logout)
		authApi.POST("/me", controllers.Me)
	}
	// router.POST("/api/auth/register", controllers.Register)
	// router.POST("/api/auth/login", controllers.Login)

	// CRUD Books
	bookApi := router.Group("/api/books")
	// bookApi.Use(middleware.RequireAuth())
	{
		bookApi.GET("/", controllers.FindBooks)
		bookApi.POST("/", controllers.StoreBook)
		bookApi.GET("/:id", controllers.FindBookById)
		bookApi.PUT("/:id", controllers.UpdateBook)
		bookApi.DELETE("/:id", controllers.DeleteBook)
	}
	// router.GET("/api/books", controllers.FindBooks)
	// router.POST("/api/books", controllers.StoreBook)
	// router.GET("/api/books/:id", controllers.FindBookById)
	// router.PUT("/api/books/:id", controllers.UpdateBook)
	// router.DELETE("/api/books/:id", controllers.DeleteBook)

	router.Run(":8080")
}
