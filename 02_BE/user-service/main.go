package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Tạo router với middleware mặc định (logging, recovery)
	router := gin.Default()

	// Định nghĩa route để lấy thông tin hồ sơ người dùng
	router.GET("/profile", handleGetUserProfile)

	// Chạy user-service trên cổng 8002
	router.Run(":8002")
}

// Handler để xử lý yêu cầu lấy thông tin hồ sơ người dùng
func handleGetUserProfile(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "User profile data",
	})
}
