package server

import (
	"chat-realtime/internal/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Server struct {
	port  int
	db    *gorm.DB
	redis *redis.Client
}

// init server
func InitServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	serverConfig := &Server{
		port:  port,
		db:    InitConnectDatabase(),
		redis: InitRedisClient(),
	}

	// Declare Server config
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", serverConfig.port),
		// Handler:      routes.InitRoutes(serverConfig.db),
		Handler:      routes.InitRoutes(serverConfig.db, serverConfig.redis),
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

func InitRedisClient() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST") // Host Redis
	redisPort := os.Getenv("REDIS_PORT") // Port Redis

	if redisHost == "" || redisPort == "" {
		log.Fatalf("Redis host or port is not set in environment variables")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})
	// Kiểm tra kết nối
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Không thể kết nối Redis: %v", err)
	}

	log.Println("Kết nối Redis thành công!")
	return client
}
