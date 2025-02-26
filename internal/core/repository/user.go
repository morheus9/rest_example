package repository

import "github.com/morheus9/rest_example/internal/core/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id int) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Update(user *entity.User) error
	Delete(id int) error
}
