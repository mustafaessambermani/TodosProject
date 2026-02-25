package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title" example:"Build My First API"`
	Description string `json:"description" example:"Using Go, Gin, and GORM"`
	Status      string `gorm:"default:'pending'" json:"status" example:"pending"`
	UserID      uint   `json:"user_id" example:"1"`
}
