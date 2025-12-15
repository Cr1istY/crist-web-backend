package service

import (
	"crist-blog/internal/repository"

	"github.com/google/uuid"
)

type CategoryService struct {
	CategoryRepo *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepo: categoryRepo,
	}
}

func (s *CategoryService) GetNameByID(id uuid.UUID) (string, error) {
	return s.CategoryRepo.GetNameByID(id)
}
