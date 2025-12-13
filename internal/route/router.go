package route

import (
	"crist-blog/internal/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	admin := e.Group("/api")
	admin.GET("/check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Admin!")
	})
	admin.POST("/login", userHandler.Login)
}
