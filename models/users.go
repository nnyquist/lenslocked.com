package models

import "gorm.io/gorm"

// User is ...
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;uniqueIndex"`
}
