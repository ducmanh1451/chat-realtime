package routes

import (
	"chat-realtime/internal/controllers"
	"chat-realtime/internal/middlewares"
	"chat-realtime/internal/repositories"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// initializes the routes for the server.
func InitRoutes(db *gorm.DB, redisClient *redis.Client) http.Handler {
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL here
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// init repository and controller
	userRepository := repositories.InitUserRepository(db)
	authRepository := repositories.InitAuthRepository(db)
	userController := controllers.InitUserController(userRepository)
	authController := controllers.InitAuthController(authRepository, redisClient)

	// declare routes
	r.POST("/login", authController.Login)
	r.POST("/logout", authController.Logout)

	validRoutes := r.Group("/protected")
	validRoutes.Use(middlewares.VerifySessionToken(redisClient)) // Apply middleware to check token
	{
		validRoutes.GET("/users", userController.GetUsers)
		validRoutes.POST("/create-user", userController.CreateUser)
	}

	return r
}
