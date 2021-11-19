package store

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(36);not null"`
	Password string `gorm:"type:varchar(128);not null"`
}

type RefreshToken struct{}
