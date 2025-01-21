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
	users, res := c.Repo.GetUsers()
	if res != nil {
		ctx.JSON(res.Status, gin.H{"status": res.Status, "message": res.Message, "payload": users})
		return
	}
}

// create user
func (c *UserController) CreateUser(ctx *gin.Context) {
	// validate body
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		bindErr := helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E003"))
		ctx.JSON(bindErr.Status, gin.H{"status": bindErr.Status, "message": bindErr.Message})
		return
	}

	// call repository
	createdUser, res := c.Repo.CreateUser(&user)
	if res != nil {
		ctx.JSON(res.Status, gin.H{"status": res.Status, "message": res.Message, "payload": createdUser})
		return
	}
}

// update user
func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var userUpdate models.User
	if err := ctx.ShouldBindJSON(&userUpdate); err != nil {
		bindErr := helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E003"))
		ctx.JSON(bindErr.Status, gin.H{"status": bindErr.Status, "message": bindErr.Message})
		return
	}

	updatedUser, res := c.Repo.UpdateUser(id, &userUpdate)
	if res != nil {
		ctx.JSON(res.Status, gin.H{"status": res.Status, "message": res.Message, "payload": updatedUser})
		return
	}
}
