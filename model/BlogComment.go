package model

import "time"

type BlogComment struct {
	ID                int        `json:"id"`
	Text              string     `json:"text"`
	UserID            int64      `json:"userId"`
	PublishedDateTime *time.Time `json:"publishedDateTime,omitempty"`
	LastUpdateTime    *time.Time `json:"lastUpdateTime,omitempty"`
}
