package handlers

import (
	"net/http"
	"pro-todo-api/internal/models"
	"pro-todo-api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Service *services.TodoService
}

func NewTodoHandler(s *services.TodoService) *TodoHandler {
	return &TodoHandler{Service: s}
}

// CreateTodo godoc
// @Summary      Create a new task
// @Description  Add a new todo item to the authenticated user's list.
// @Description  Note: status can be 'pending', 'in-progress', or 'completed'.
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        todo  body      models.Todo  true  "Todo Data (Example: {'title': 'Task Name', 'description': 'Details'})"
// @Success      201   {object}  models.Todo
// @Security     BearerAuth
// @Router       /todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real scenario, we get UserID from JWT middleware
	userID := c.MustGet("user_id").(uint)
	todo.UserID = userID

	if err := h.Service.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetTodos godoc
// @Summary      List user todos
// @Description  Get all tasks for the authenticated user
// @Tags         Todos
// @Produce      json
// @Security     BearerAuth
// @Success      200   {array}   models.Todo
// @Router       /todos [get]
func (h *TodoHandler) GetTodos(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	todos, err := h.Service.GetUserTodos(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// UpdateTodo godoc
// @Summary      Update a todo
// @Description  Update title, description or status of a specific todo
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int          true  "Todo ID"
// @Param        todo  body      models.Todo  true  "Updated Data"
// @Success      200   {object}  map[string]string
// @Router       /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("user_id").(uint)

	var updatedData models.Todo
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateTodo(uint(id), userID, updatedData); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo updated successfully"})
}

// DeleteTodo godoc
// @Summary      Delete a todo
// @Description  Remove a task by ID
// @Tags         Todos
// @Security     BearerAuth
// @Param        id   path      int  true  "Todo ID"
// @Success      204   "No Content"
// @Router       /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("user_id").(uint)

	if err := h.Service.DeleteTodo(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
