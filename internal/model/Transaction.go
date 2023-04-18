package model

import (
	"time"

	"src/gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID          uint64    `gorm:"primary_key;auto_increment"`
	Date        time.Time `json:"date"`
	Transaction float64   `json:"transaction"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}