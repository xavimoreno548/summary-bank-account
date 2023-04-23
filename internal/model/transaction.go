package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Date        time.Time `json:"date"`
	Transaction float64   `json:"transaction"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
