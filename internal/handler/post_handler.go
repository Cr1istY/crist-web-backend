package handler

import (
	"crist-blog/internal/model"
	"crist-blog/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

type CreatePostRequest struct {
	UserID          string     `json:"user_id" validate:"required,uuid4"`
	Title           string     `json:"title" validate:"required"`
	Slug            string     `json:"slug" validate:"required"`
	Content         string     `json:"content"`
	Excerpt         string     `json:"excerpt"`
	Status          string     `json:"status" validate:"oneof=draft published private"`
	CategoryID      string     `json:"category_id" validate:"required,uuid4"`
	Tags            []string   `json:"tags"`
	MetaTitle       string     `json:"meta_title"`
	MetaDescription string     `json:"meta_description"`
	PublishedAt     *time.Time `json:"published_at"`
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	var req CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if _, err := uuid.Parse(req.UserID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	if _, err := uuid.Parse(req.CategoryID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid category ID"})
	}

	post := &model.Post{
		UserID:          uuid.MustParse(req.UserID),
		Title:           req.Title,
		Slug:            req.Slug,
		Content:         req.Content,
		Excerpt:         req.Excerpt,
		Status:          model.PostStatus(req.Status),
		CategoryID:      uuid.MustParse(req.CategoryID),
		Tags:            req.Tags,
		MetaTitle:       req.MetaTitle,
		MetaDescription: req.MetaDescription,
		PublishedAt:     req.PublishedAt,
	}
	if err := h.postService.CreatePost(post); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, post)
}

func (h *PostHandler) Get(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	id := uint(id64)
	post, err := h.postService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, post)
}

func (h *PostHandler) Update(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	id := uint(id64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var req CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	post := &model.Post{
		ID:              id,
		UserID:          uuid.MustParse(req.UserID), // 注意：实际应从 token 获取，不应由前端传
		Title:           req.Title,
		Slug:            req.Slug,
		Content:         req.Content,
		Excerpt:         req.Excerpt,
		Status:          model.PostStatus(req.Status),
		CategoryID:      uuid.MustParse(req.CategoryID),
		Tags:            req.Tags,
		MetaTitle:       req.MetaTitle,
		MetaDescription: req.MetaDescription,
		PublishedAt:     req.PublishedAt,
	}

	if err := h.postService.Update(post); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	updated, _ := h.postService.GetByID(id)
	return c.JSON(http.StatusOK, updated)
}

func (h *PostHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	id := uint(id64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.postService.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *PostHandler) List(c echo.Context) error {
	posts, err := h.postService.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, posts)
}
