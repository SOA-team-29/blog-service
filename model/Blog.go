package model

import (
	"time"

	"github.com/lib/pq"
)

type BlogComments []BlogComment
type BlogRatings []BlogRating
type BlogPostStatus int

const (
	DRAFT BlogPostStatus = iota
	PUBLISHED
	CLOSED
	ACTIVE
	FAMOUS
)

type Blog struct {
	ID           int            `json:"id"`
	AuthorID     int            `json:"authorId"`
	TourID       int            `json:"tourId"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	CreationDate *time.Time     `json:"creationDate,omitempty"`
	ImageURLs    pq.StringArray `json:"imageURLs" gorm:"type:text[]"`
	Comments     BlogComments   `json:"comments" gorm:"type:jsonb"`
	Ratings      BlogRatings    `json:"ratings" gorm:"type:jsonb"`
	Status       BlogPostStatus `json:"status"`
}
