package controllers

import (
	"backend/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidateBookInput struct {
	Name        string `json:"name" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Publisher   string `json:"publisher" binding:"required"`
	PublishDate int    `json:"publish_date" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "unkown error"
}

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Lists Data Books",
		"data":    books,
	})
}

func StoreBook(c *gin.Context) {
	var input ValidateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	book := models.Book{
		Name:        input.Name,
		Author:      input.Author,
		Publisher:   input.Publisher,
		PublishDate: input.PublishDate,
		Price:       input.Price,
		Stock:       input.Stock,
	}
	models.DB.Create(&book)

	c.JSON(201, gin.H{
		"success": true,
		"message": "Book Created Successfully",
		"data":    book,
	})
}

func FindBookById(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Post By ID : " + c.Param("id"),
		"data":    book,
	})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input ValidateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	models.DB.Model(&book).Updates(input)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Book Updated Successfully",
		"data":    book,
	})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Book deleted successfully",
	})
}
