package repository

import (
	"log"
	"pro-todo-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBRawURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
