package models

import (
	"time"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	GoogleSub  string `gorm:"uniqueIndex"` // Google用户唯一标识
	Email      string `gorm:"uniqueIndex"`
	Name       string
	PictureURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
