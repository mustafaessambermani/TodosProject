package services

import (
	"errors"
	"pro-todo-api/internal/models"
	"pro-todo-api/internal/repository"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

func NewTodoService(r *repository.TodoRepository) *TodoService {
	return &TodoService{Repo: r}
}

// CreateTodo validates and saves a new task
func (s *TodoService) CreateTodo(todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("todo title cannot be empty")
	}
	return s.Repo.Create(todo)
}

// GetUserTodos fetches all tasks for a specific user
func (s *TodoService) GetUserTodos(userID uint) ([]models.Todo, error) {
	return s.Repo.GetAllByUserID(userID)
}

// UpdateTodoStatus updates the task details if the user owns it
func (s *TodoService) UpdateTodo(id uint, userID uint, updatedData models.Todo) error {
	// 1. Fetch the existing todo to verify ownership
	todo, err := s.Repo.GetByID(id)
	if err != nil {
		return errors.New("todo not found")
	}

	if todo.UserID != userID {
		return errors.New("unauthorized: you do not own this todo")
	}

	// 2. Update fields
	todo.Title = updatedData.Title
	todo.Description = updatedData.Description
	todo.Status = updatedData.Status

	return s.Repo.Update(todo)
}

// DeleteTodo removes a task after security verification
func (s *TodoService) DeleteTodo(id uint, userID uint) error {
	return s.Repo.Delete(id, userID)
}
