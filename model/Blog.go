package model

import (
	"time"

	"github.com/lib/pq"
)

type BlogPostStatus int

const (
	Draft BlogPostStatus = iota
	Published
	Closed
)

type Blog struct {
	AuthorID            int64          `json:"aid"`
	TourID              int64          `json:"tid"`
	Title               string         `json:"title"`
	Description         string         `json:"description"`
	DescriptionMarkdown string         `json:"descriptionMarkdown"`
	PublishedDateTime   *time.Time     `json:"publishedDateTime,omitempty"`
	ImageID             pq.StringArray `json:"images" gorm:"type:text[]"`
	Status              BlogPostStatus `json:"status"`
	BlogComments        []BlogComment  `json:"blogCom" gorm:"type:json"`
	BlogRatings         []BlogRating   `json:"blogRate" gorm:"type:json"`
}
