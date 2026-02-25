package repository

import (
	"pro-todo-api/internal/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

// Create persists a new todo record in the database
func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.DB.Create(todo).Error
}

// GetAllByUserID retrieves all todo tasks belonging to a specific user
func (r *TodoRepository) GetAllByUserID(userID uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.DB.Where("user_id = ?", userID).Find(&todos).Error
	return todos, err
}

// GetByID retrieves a single todo task by its primary key
func (r *TodoRepository) GetByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.DB.First(&todo, id).Error
	return &todo, err
}

// Update modifies an existing todo record
func (r *TodoRepository) Update(todo *models.Todo) error {
	return r.DB.Save(todo).Error
}

// Delete removes a todo record only if it belongs to the authenticated user
func (r *TodoRepository) Delete(id uint, userID uint) error {
	// Security check: Ensure the todo belongs to the user before deleting
	return r.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{}).Error
}
