package models

import "gorm.io/gorm"

// User represents the user model in the system
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" example:"qaseh_dev"`
	Email    string `gorm:"unique;not null" json:"email" example:"qaseh@example.com"`
	Password string `gorm:"not null" json:"password" example:"secret123"`
}
