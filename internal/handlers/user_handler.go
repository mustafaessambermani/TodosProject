package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-todo-api/internal/dtos"
	"pro-todo-api/internal/services"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// SignUp godoc
// @Summary      User Registration
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      dtos.SignUpRequest  true  "User Data"
// @Success      201   {object}  map[string]string
// @Router       /auth/signup [post]
func (h *UserHandler) SignUp(c *gin.Context) {
	var req dtos.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.SignUp(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

// Login godoc
// @Summary      User Login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      dtos.LoginRequest  true  "Credentials"
// @Success      200   {object}  map[string]string
// @Router       /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dtos.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
