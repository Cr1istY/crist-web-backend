package route

import (
	"crist-blog/internal/handler"

	"github.com/labstack/echo/v4"
)

func SetupBlogRouter(e *echo.Echo, postHandler *handler.PostHandler) {
	api := e.Group("/api")
	posts := api.Group("/posts")
	posts.POST("/create", postHandler.CreatePost)
	posts.GET("/getAllPosts", postHandler.ListToFrontend)
	posts.GET("/get/:id", postHandler.Get)
	posts.PUT("/update/:id", postHandler.Update)
	posts.DELETE("/delete/:id", postHandler.Delete)
}
