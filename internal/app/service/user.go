package service

import (
	"github.com/morheus9/rest_example/internal/core/entity"
	"github.com/morheus9/rest_example/internal/core/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) (*entity.User, error) {
	user := &entity.User{
		Name:  name,
		Email: email,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// Реализация остальных методов...
