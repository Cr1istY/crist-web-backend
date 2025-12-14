package repository

import (
	"crist-blog/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(post *model.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) GetByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.Where("id = ?", id).First(&post).Error
	return &post, err
}

func (r *PostRepository) Update(post *model.Post) error {
	return r.DB.Save(post).Error
}

func (r *PostRepository) Delete(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&model.Post{}).Error
}

func (r *PostRepository) List() ([]*model.Post, error) {
	var posts []*model.Post
	return posts, r.DB.Find(&posts).Error
}
