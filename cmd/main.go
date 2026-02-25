package main

import (
	"log"
	"pro-todo-api/config"
	"pro-todo-api/internal/handlers"
	"pro-todo-api/internal/middleware"
	"pro-todo-api/internal/models"
	"pro-todo-api/internal/repository"
	"pro-todo-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	// This is important for Swagger
	_ "pro-todo-api/docs"
)

// @title           Pro Todo API
// @version         1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// 1. Setup Configuration
	cfg := config.LoadConfig()

	// 2. Database Initialization & Auto Migration
	db := repository.InitDB(cfg)
	db.AutoMigrate(&models.User{}, &models.Todo{})

	// 3. Dependency Injection (Wiring the layers)
	userRepo := repository.NewUserRepository(db)
	todoRepo := repository.NewTodoRepository(db)

	userService := services.NewUserService(userRepo, cfg)
	todoService := services.NewTodoService(todoRepo)

	userHandler := handlers.NewUserHandler(userService)
	todoHandler := handlers.NewTodoHandler(todoService)

	// 4. Router Setup
	r := gin.Default()

	// Swagger Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth Routes (Public)
	auth := r.Group("/auth")
	{
		auth.POST("/signup", userHandler.SignUp)
		auth.POST("/login", userHandler.Login)
	}

	// Todo Routes (Protected by Middleware)
	todoRoutes := r.Group("/todos")
	todoRoutes.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		todoRoutes.POST("/", todoHandler.CreateTodo)
		todoRoutes.GET("/", todoHandler.GetTodos)
		todoRoutes.PUT("/:id", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodo)
	}

	// 5. Start Server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
