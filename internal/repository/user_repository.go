package repository

import (
	"gorm.io/gorm"
	"pro-todo-api/internal/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
