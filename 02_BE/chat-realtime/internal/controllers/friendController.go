package controllers

import (
	"chat-realtime/internal/helpers"
	"chat-realtime/internal/repositories"

	"github.com/gin-gonic/gin"
)

type FriendController struct {
	Repo *repositories.FriendRepository
}

// init controller
func InitFriendController(repo *repositories.FriendRepository) *FriendController {
	return &FriendController{Repo: repo}
}

// send friend request
func (c *FriendController) SendFriendRequest(ctx *gin.Context) {
	var params struct {
		SenderID   int    `json:"sender_id"`
		SenderName string `json:"sender_name"`
		ReceiverID int    `json:"receiver_id"`
	}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		bindErr := helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E003"))
		ctx.JSON(bindErr.Status, gin.H{"status": bindErr.Status, "message": bindErr.Message})
		return
	}

	_, res := c.Repo.SendFriendRequest(params.SenderID, params.SenderName, params.ReceiverID)
	if res != nil {
		ctx.JSON(res.Status, gin.H{"status": res.Status, "message": res.Message})
		return
	}
}

// accept friend request
func (c *FriendController) AcceptFriendRequest(ctx *gin.Context) {
	var params struct {
		SenderID     int    `json:"sender_id"`
		ReceiverID   int    `json:"receiver_id"`
		ReceiverName string `json:"receiver_name"`
	}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		bindErr := helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E003"))
		ctx.JSON(bindErr.Status, gin.H{"status": bindErr.Status, "message": bindErr.Message})
		return
	}

	_, res := c.Repo.AcceptFriendRequest(params.SenderID, params.ReceiverID, params.ReceiverName)
	if res != nil {
		ctx.JSON(res.Status, gin.H{"status": res.Status, "message": res.Message})
		return
	}
}

// delete friend request
func (c *FriendController) DeleteFriendRequest(ctx *gin.Context) {
	var params struct {
		SenderID   int `json:"sender_id"`
		ReceiverID int `json:"receiver_id"`
	}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		bindErr := helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E003"))
		ctx.JSON(bindErr.Status, gin.H{"status": bindErr.Status, "message": bindErr.Message})
		return
	}

	_, res := c.Repo.DeleteFriendRequest(params.SenderID, params.ReceiverID)
	if res != nil {
		ctx.JSON(res.Status, gin.H{"status": res.Status, "message": res.Message})
		return
	}
}
