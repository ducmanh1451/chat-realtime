package controllers

import (
	"chat-realtime/internal/helpers"
	"chat-realtime/internal/models"
	"chat-realtime/internal/repositories"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo *repositories.UserRepository
}

// init controller
func InitUserController(repo *repositories.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

// get users
func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.Repo.GetUsers()
	if err != nil {
		ctx.JSON(err.Status, gin.H{"error": gin.H{
			"status":  err.Status,
			"message": err.Message,
		}})
		return
	}
	ctx.JSON(helpers.StatusOK, users)
}

// create user
func (c *UserController) CreateUser(ctx *gin.Context) {
	// validate body
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		bindErr := helpers.CreateError(helpers.StatusBadRequest, helpers.GetErrorMessage("E003"))
		ctx.JSON(bindErr.Status, gin.H{"error": gin.H{
			"status":  bindErr.Status,
			"message": bindErr.Message,
		}})
		return
	}

	// call repository
	createdUser, err := c.Repo.CreateUser(&user)
	if err != nil {
		ctx.JSON(err.Status, gin.H{"error": gin.H{
			"status":  err.Status,
			"message": err.Message,
		}})
		return
	}

	// return data
	ctx.JSON(helpers.StatusOK, createdUser)
}
