package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// StatsJSON represents JSON daily statistics entry
type StatsJSON struct {
	CreatedAt time.Time `json:"created_at" validation:"required"`
	Category  int       `json:"category" validation:"required"`
	Data      DataMap   `json:"data" validation:"required"`
}

type DataMap map[string]interface{}

func (a DataMap) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// StatsFP represents floating point daily statistics entry
type StatsFP struct {
	CreatedAt time.Time `json:"created_at" validation:"required"`
	Category  int       `json:"category" validation:"required"`
	Data      float64   `json:"data" validation:"required"`
}

// StatsInt represents integer statistics entry
type StatsInt struct {
	CreatedAt time.Time `json:"created_at" validation:"required"`
	Category  int       `json:"category" validation:"required"`
	Data      int       `json:"data" validation:"required"`
}
