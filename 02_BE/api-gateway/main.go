package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Tạo một router với middleware mặc định (logging, recovery)
	router := gin.Default()

	// TEST
	router.Any("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "TEST 5",
		})
	})

	// Định nghĩa route để chuyển tiếp yêu cầu tới auth-service
	router.Any("/auth/*authServicePath", func(context *gin.Context) {
		fmt.Println("Request Method:", context)
		// Đây là code giả lập để chuyển tiếp yêu cầu đến auth-service
		context.JSON(http.StatusOK, gin.H{
			"message": "Forwarding request to auth-service aaaaaa bbbb",
		})
	})

	// Định nghĩa route để chuyển tiếp yêu cầu tới user-service
	router.Any("/user/*userServicePath", func(context *gin.Context) {
		fmt.Println("Request Method:", context)
		// Đây là code giả lập để chuyển tiếp yêu cầu đến user-service
		context.JSON(http.StatusOK, gin.H{
			"message": "Forwarding request to user-service",
		})
	})

	// Chạy api-gateway trên cổng 8000
	router.Run(":8000")
}
