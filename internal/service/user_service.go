package service

import (
	"context"
	"errors"

	"github.com/morheus9/rest_example/internal/domain"
	"github.com/morheus9/rest_example/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id int64) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if user.Name == "" || user.Email == "" {
		return nil, errors.New("name and email cannot be empty")
	}
	return s.repo.Create(ctx, user)
}

func (s *userService) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}
