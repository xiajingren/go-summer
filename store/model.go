package store

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(64);not null"`
	Password string `gorm:"type:varchar(256);not null"`
}

type RefreshToken struct {
}
