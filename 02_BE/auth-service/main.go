package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Tạo router với middleware mặc định (logging, recovery)
	router := gin.Default()

	// Định nghĩa route cho đăng ký người dùng
	router.POST("/register", handleUserRegister)

	// Định nghĩa route cho đăng nhập người dùng
	router.POST("/login", handleUserLogin)

	// Chạy auth-service trên cổng 8001
	router.Run(":8001")
}

// Handler để xử lý đăng ký người dùng
func handleUserRegister(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "User registered successfully",
	})
}

// Handler để xử lý đăng nhập người dùng
func handleUserLogin(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "User logged in successfully",
	})
}
