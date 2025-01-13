package server

import (
	"chat-realtime/internal/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Server struct {
	port int
	db   *gorm.DB
}

// init server
func InitServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	serverConfig := &Server{
		port: port,
		db:   InitConnectDatabase(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", serverConfig.port),
		Handler:      routes.InitRoutes(serverConfig.db),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

// init connection
func InitConnectDatabase() *gorm.DB {
	dbname := os.Getenv("BLUEPRINT_DB_DATABASE")
	password := os.Getenv("BLUEPRINT_DB_PASSWORD")
	username := os.Getenv("BLUEPRINT_DB_USERNAME")
	portDB := os.Getenv("BLUEPRINT_DB_PORT")
	host := os.Getenv("BLUEPRINT_DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, portDB, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Optional: Log SQL queries for debugging
	})

	if err != nil {
		log.Fatalf("Lỗi kết nối cơ sở dữ liệu: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Không thể cấu hình nhóm kết nối cơ sở dữ liệu: %v", err)
	}

	// Configure connection pool settings
	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(50)

	return db
}
