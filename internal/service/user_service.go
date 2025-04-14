package service

import (
	"errors"
	"strings"
	"time"

	"github.com/Savitree1999/app-service/internal/repository"
	"github.com/Savitree1999/app-service/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ErrEmailExists = errors.New("Email already exists")

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	repo *repository.UserRepository
	log  *zap.SugaredLogger
}

func NewUserService(db *gorm.DB, log *zap.SugaredLogger) *UserService {
	repo := repository.NewUserRepository(db, log)
	return &UserService{
		repo: repo,
		log:  log,
	}
}

func (s *UserService) CreateUser(req CreateUserRequest) (*model.User, error) {
	user := &model.User{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	// Log request details
	s.log.Infof("Received request to create user: %v", req)

	// Check if email already exists
	if err := s.repo.Create(user); err != nil {
		s.log.Errorf("Error creating user: %v", err)
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, ErrEmailExists
		}
		return nil, err
	}

	// Log successful user creation
	s.log.Infof("User created successfully: %v", user)
	return user, nil
}
