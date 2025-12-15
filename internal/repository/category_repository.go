package repository

import (
	"crist-blog/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetNameByID(id uuid.UUID) (string, error) {
	var name string
	err := r.DB.Model(&model.Category{}).
		Select("name").
		Where("id = ?", id).
		First(&name).Error
	return name, err
}
