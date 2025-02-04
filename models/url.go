package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	LongUrl    string    `json:"long_url" gorm:"not null"`
	ShortUrl   string    `json:"short_url" gorm:"unique;not null"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
	ClickCount int       `json:"default:0"`
}
