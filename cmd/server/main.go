package main

import (
	"crist-blog/internal/blogConfig"
	"crist-blog/internal/handler"
	"crist-blog/internal/repository"
	"crist-blog/internal/route"
	"crist-blog/internal/service"
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

// 使用示例
func main() {

	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)
	if err != nil {
	}
	jwtSecret := base64.URLEncoding.EncodeToString(bytes)

	db := blogConfig.ConnectDB()
	userRepo := repository.NewUserRepository(db)
	authRepo := repository.NewRefreshTokenRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, authRepo, jwtSecret)
	userHandler := handler.NewUserHandler(authService, userService)

	e := echo.New()
	route.SetupUserRoutes(e, userHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server is running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
