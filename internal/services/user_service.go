package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"pro-todo-api/config"
	"pro-todo-api/internal/dtos"
	"pro-todo-api/internal/models"
	"pro-todo-api/internal/repository"
	"pro-todo-api/pkg/utils"
)

type UserService struct {
	Repo   *repository.UserRepository
	Config *config.Config
}

func NewUserService(r *repository.UserRepository, cfg *config.Config) *UserService {
	return &UserService{Repo: r, Config: cfg}
}

func (s *UserService) SignUp(req dtos.SignUpRequest) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	return s.Repo.CreateUser(&user)
}

func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.Repo.GetByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID, s.Config.JWTSecret)
}
