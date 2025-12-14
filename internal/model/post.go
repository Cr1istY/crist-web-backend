package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type PostStatus string

const (
	Draft     PostStatus = "draft"
	Published PostStatus = "published"
)

type Post struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"id"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	Title           string         `gorm:"type:text;not null" json:"title"`
	Slug            string         `gorm:"type:text;not null;uniqueIndex" json:"slug"`
	Content         string         `gorm:"type:text" json:"content"`
	Excerpt         string         `gorm:"type:text" json:"excerpt"`
	Status          PostStatus     `gorm:"type:post_status_enum;not null" json:"status"`
	CategoryID      uuid.UUID      `gorm:"type:uuid;not null" json:"category_id"`
	Tags            pq.StringArray `gorm:"type:text[]" json:"tags"` // 👈 关键修改
	Views           int            `gorm:"default:0" json:"views"`
	Likes           int            `gorm:"default:0" json:"likes"`
	PublishedAt     *time.Time     `json:"published_at"`
	MetaTitle       string         `gorm:"type:text" json:"meta_title"`
	MetaDescription string         `gorm:"type:text" json:"meta_description"`
	SearchVector    interface{}    `gorm:"type:tsvector" json:"-"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Post) TableName() string {
	return "blog.posts"
}
