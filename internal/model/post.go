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
	Private   PostStatus = "private"
)

type Post struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	Title           string         `gorm:"type:text;not null" json:"title"`
	Slug            string         `gorm:"type:text;not null;uniqueIndex" json:"slug"`
	Content         string         `gorm:"type:text" json:"content"`
	Excerpt         string         `gorm:"type:text" json:"excerpt"`
	Status          PostStatus     `gorm:"type:post_status_enum;not null" json:"status"`
	CategoryID      uuid.UUID      `gorm:"type:uuid;not null" json:"category_id"`
	Tags            pq.StringArray `gorm:"type:text[]" json:"tags"`
	Views           int            `gorm:"default:0" json:"views"`
	Likes           int            `gorm:"default:0" json:"likes"`
	Thumbnail       string         `gorm:"type:text" json:"thumbnail"`
	PublishedAt     *time.Time     `json:"published_at"`
	MetaTitle       string         `gorm:"type:text" json:"meta_title"`
	MetaDescription string         `gorm:"type:text" json:"meta_description"`
	SearchVector    interface{}    `gorm:"type:tsvector" json:"-"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type PostFrontend struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	Date      string   `json:"date"`
	Excerpt   string   `json:"excerpt"`
	Views     int      `gorm:"default:0" json:"views"`
	Likes     int      `gorm:"default:0" json:"likes"`
	Thumbnail string   `json:"thumbnail,omitempty"`
}

type HotPost struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string    `gorm:"type:text;not null" json:"title"`
	CategoryID uuid.UUID `gorm:"type:uuid;not null" json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	Excerpt    string    `gorm:"type:text" json:"excerpt"`
}

type HotPostFrontend struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title    string `gorm:"type:text;not null" json:"title"`
	Category string `json:"category"`
	Date     string `json:"date"`
	Excerpt  string `gorm:"type:text" json:"excerpt"`
}

type LatestPost struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string    `gorm:"type:text;not null" json:"title"`
	CategoryID uuid.UUID `gorm:"type:uuid;not null" json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type LatestPostFrontend struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title    string `gorm:"type:text;not null" json:"title"`
	Date     string `json:"date"`
	Category string `json:"category"`
}

func (Post) TableName() string {
	return "blog.posts"
}
