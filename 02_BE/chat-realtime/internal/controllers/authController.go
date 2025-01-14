package controllers

import (
	"chat-realtime/internal/helpers"
	"chat-realtime/internal/repositories"
	"context"
	"log"
	"time"

	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type AuthController struct {
	Repo  *repositories.AuthRepository
	Redis *redis.Client
}

// init controller
func InitAuthController(repo *repositories.AuthRepository, redisClient *redis.Client) *AuthController {
	return &AuthController{Repo: repo, Redis: redisClient}
}

// login
func (c *AuthController) Login(ctx *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		bindErr := helpers.CreateError(helpers.StatusBadRequest, helpers.GetErrorMessage("E003"))
		ctx.JSON(bindErr.Status, gin.H{
			"error": gin.H{
				"status":  bindErr.Status,
				"message": bindErr.Message,
			},
		})
		return
	}
	user, err := c.Repo.Login(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(err.Status, gin.H{
			"error": gin.H{
				"status":  err.Status,
				"message": err.Message,
			},
		})
		return
	}

	// create redis session token
	sessionToken := GenerateRandomString(32)
	// save redis session token with a 24-hour expiration time
	ctxRedis := context.Background()
	redisErr := c.Redis.Set(ctxRedis, sessionToken, user.ID, 24*time.Hour).Err()
	if redisErr != nil {
		redisHelperErr := helpers.CreateError(helpers.StatusBadRequest, helpers.GetErrorMessage("E005"))
		ctx.JSON(redisHelperErr.Status, gin.H{
			"error": gin.H{
				"status":  redisHelperErr.Status,
				"message": redisHelperErr.Message,
			},
		})
		return
	}
	ctx.JSON(helpers.StatusOK, gin.H{
		"user":  user,
		"token": sessionToken,
	})
}

// logout
func (c *AuthController) Logout(ctx *gin.Context) {
	sessionToken := ctx.GetHeader("Authorization")

	if sessionToken == "" {
		err := helpers.CreateError(helpers.StatusBadRequest, helpers.GetErrorMessage("E006"))
		ctx.JSON(err.Status, gin.H{
			"error": gin.H{
				"status":  err.Status,
				"message": err.Message,
			},
		})
		return
	}

	ctxRedis := context.Background()
	err := c.Redis.Del(ctxRedis, sessionToken).Err()
	if err != nil {
		logoutErr := helpers.CreateError(helpers.StatusBadRequest, helpers.GetErrorMessage("E007"))
		ctx.JSON(logoutErr.Status, gin.H{
			"error": gin.H{
				"status":  logoutErr.Status,
				"message": logoutErr.Message,
			},
		})
		return
	}
	ctx.JSON(helpers.StatusOK, nil)
}

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatalf("Không thể tạo chuỗi ngẫu nhiên: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes) // Sử dụng base64 để mã hóa thành chuỗi
}
