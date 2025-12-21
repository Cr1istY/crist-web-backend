package handler

import (
	"crist-blog/internal/model"
	"crist-blog/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) ListAllCategories(c echo.Context) error {
	var categories []model.CreatePostCategory
	rawCategories, err := h.categoryService.ListAllCategories()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	for _, rawCategory := range rawCategories {
		categories = append(categories, model.CreatePostCategory{
			ID:   rawCategory.ID,
			Name: rawCategory.Name,
		})
	}
	return c.JSON(http.StatusOK, categories)
}
