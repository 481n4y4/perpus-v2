package controllers

import (
	"backend/models"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type ValidateRegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ValidateLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input ValidateRegisterInput
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

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash),
	}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User register successfully",
		"data":    user,
	})
}

func Login(c *gin.Context) {
	var input ValidateLoginInput
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

	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create token",
		})
		return
	}

	c.SetCookie(
		"auth_token",
		tokenString,
		3600*24,
		"/",
		"",
		false,
		false,
	)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successfully",
		"token":   tokenString,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie(
		"auth_token", // Nama cookie yang mau dihapus
		"",           // Isinya dikosongkan
		-1,           // MaxAge -1 memaksa browser untuk langsung menghapus cookie ini detik ini juga
		"/",          // Path harus sama dengan saat kamu create cookie di fungsi Login
		"",           // Domain disamakan
		false,        // Secure disamakan
		false,        // HttpOnly false
	)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout successfully",
	})
}

func Me(c *gin.Context) {
	tokenString, err := c.Cookie("auth_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Kamu belum login atau sesi telah berakhir",
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("method signing token tidak valid")
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token tidak valid atau sudah kedaluwarsa",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Gagal membaca data di dalam token",
		})
	}

	userId := claims["sub"]

	var user models.User
	if err := models.DB.First(&user, userId).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User tidak ditemukan di sistem",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil mengambil data profile",
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
