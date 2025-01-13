package routes

import (
	"chat-realtime/internal/controllers"
	"chat-realtime/internal/repositories"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// initializes the routes for the server.
func InitRoutes(db *gorm.DB) http.Handler {
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
	userController := controllers.InitUserController(userRepository)

	// declare routes
	r.GET("/users", userController.GetUsers)
	r.POST("/create-user", userController.CreateUser)

	return r
}
