package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type BlogComment struct {
	Text            string    `json:"text"`
	UserID          int       `json:"userId"`
	CreationTime    time.Time `json:"creationTime,omitempty"`
	LastUpdatedTime time.Time `json:"lastUpdatedTime,omitempty"`
}

func (t *BlogComment) Scan(value interface{}) error {
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
func (t BlogComment) Value() (driver.Value, error) {
	return json.Marshal(t)
}
