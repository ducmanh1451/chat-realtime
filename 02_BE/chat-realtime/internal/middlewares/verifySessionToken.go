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
			unauthorizedErr := helpers.CreateError(helpers.StatusBadRequest, helpers.GetErrorMessage("E006"))
			ctx.JSON(unauthorizedErr.Status, gin.H{
				"error": gin.H{
					"status":  unauthorizedErr.Status,
					"message": unauthorizedErr.Message,
				},
			})
			ctx.Abort()
			return
		}

		// check token in redis
		ctxRedis := context.Background()
		_, err := redisClient.Get(ctxRedis, sessionToken).Result()
		if err == redis.Nil {
			unauthorizedErr := helpers.CreateError(helpers.StatusUnauthorized, helpers.GetErrorMessage("E008"))
			ctx.JSON(unauthorizedErr.Status, gin.H{
				"error": gin.H{
					"status":  unauthorizedErr.Status,
					"message": unauthorizedErr.Message,
				},
			})
			ctx.Abort()
			return
		} else if err != nil {
			// server error
			unauthorizedErr := helpers.CreateError(helpers.StatusInternalServerError, helpers.GetErrorMessage("E001"))
			ctx.JSON(unauthorizedErr.Status, gin.H{
				"error": gin.H{
					"status":  unauthorizedErr.Status,
					"message": unauthorizedErr.Message,
				},
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
