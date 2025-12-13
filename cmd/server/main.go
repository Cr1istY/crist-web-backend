package main

import (
	"crist-blog/internal/blogConfig"
	"crist-blog/internal/handler"
	"crist-blog/internal/repository"
	"crist-blog/internal/route"
	"crist-blog/internal/service"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

// 使用示例
func main() {
	db := blogConfig.ConnectDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	e := echo.New()
	route.SetupUserRoutes(e, userHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server is running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
