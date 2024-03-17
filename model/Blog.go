package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type BlogPostStatus int

const (
	Draft BlogPostStatus = iota
	Published
	Closed
)

type Blog struct {
	ID                uuid.UUID      `json:"id"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	PublishedDateTime *time.Time     `json:"publishedDateTime,omitempty"`
	ImageID           pq.StringArray `json:"images" gorm:"type:text[]"`
	Status            BlogPostStatus `json:"status"`
	UserID            int            `json:"userId"`
}

func (blog *Blog) BeforeCreate(scope *gorm.DB) error {
	blog.ID = uuid.New()
	return nil
}
