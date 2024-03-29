package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type BlogRating struct {
	IsPositive   bool       `json:"isPositive"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	UserID       int        `json:"userId"`
}

func (t *BlogRating) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, t)
}

// Value implements the driver.Valuer interface
func (t BlogRating) Value() (driver.Value, error) {
	return json.Marshal(t)
}
