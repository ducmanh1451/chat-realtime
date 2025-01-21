package middlewares

import (
	"chat-realtime/internal/helpers"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// VerifySessionToken middleware to check if the session token is valid in Redis
func VerifySessionToken(redisClient *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionToken := ctx.GetHeader("Authorization")
		if sessionToken == "" {
			unauthorizedErr := helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E006"))
			ctx.JSON(unauthorizedErr.Status, gin.H{"status": unauthorizedErr.Status, "message": unauthorizedErr.Message})
			ctx.Abort()
			return
		}

		// check token in redis
		ctxRedis := context.Background()
		_, err := redisClient.Get(ctxRedis, sessionToken).Result()
		if err == redis.Nil {
			unauthorizedErr := helpers.CreateResponse(helpers.StatusUnauthorized, helpers.GetMessage("E008"))
			ctx.JSON(unauthorizedErr.Status, gin.H{"status": unauthorizedErr.Status, "message": unauthorizedErr.Message})
			ctx.Abort()
			return
		} else if err != nil {
			// server error
			unauthorizedErr := helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
			ctx.JSON(unauthorizedErr.Status, gin.H{"status": unauthorizedErr.Status, "message": unauthorizedErr.Message})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
