package repository

import (
	"context"
	"database/sql"

	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
	"github.com/mauromamani/go-clean-architecture/internal/user/entity"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &userRepository{
		DB: db,
	}
}

// GetUser:
func (r *userRepository) GetUser(ctx context.Context) (*entity.User, error) {
	return nil, nil
}

// GetUserById:
func (r *userRepository) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	return nil, nil
}

// GetUserByEmail:
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

// CreateUser:
func (r *userRepository) CreateUser(ctx context.Context, user *dtos.CreateUserDto) (*entity.User, error) {
	return nil, nil
}

// UpdateUser:
func (r *userRepository) UpdateUser(ctx context.Context, id int, user *dtos.UpdateUserDto) (*entity.User, error) {
	return nil, nil
}

// DeleteUser:
func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	return nil
}
