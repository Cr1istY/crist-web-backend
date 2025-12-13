package repository

import (
	"crist-blog/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
func (r *UserRepository) GetByName(name string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", name).First(&user).Error
	return &user, err
}
