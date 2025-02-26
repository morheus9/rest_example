package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/morheus9/rest_example/internal/core/entity"
	"github.com/morheus9/rest_example/internal/core/repository"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) repository.UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *entity.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	return r.db.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt)
}

// Реализация остальных методов...
