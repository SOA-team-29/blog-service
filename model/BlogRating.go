package model

import "time"

type BlogRating struct {
	ID               int        `json:"id"`
	UserID           int64      `json:"userId"`
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
	IsPositive       bool       `json:"isPositive"`
}
