package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/lib/pq"
)

type BlogPostStatus int

const (
	Draft BlogPostStatus = iota
	Published
	Closed
)

type Blog struct {
	AuthorID          uuid.UUID      `json:"aid"`
	TourID            uuid.UUID      `json:"tid"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	PublishedDateTime *time.Time     `json:"publishedDateTime,omitempty"`
	ImageID           pq.StringArray `json:"images" gorm:"type:text[]"`
	Status            BlogPostStatus `json:"status"`
	BlogComments      []BlogComment  `json:"blogCom" gorm:"type:json"`
	BlogRatings       []BlogRating   `json:"blogRate" gorm:"type:json"`
}
