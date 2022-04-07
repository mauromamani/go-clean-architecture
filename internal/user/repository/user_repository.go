package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

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
func (r *userRepository) GetUsers(ctx context.Context) ([]*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, getUsersQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*entity.User{}
	for rows.Next() {
		var u entity.User
		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserById:
func (r *userRepository) GetUserById(ctx context.Context, id int64) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var u entity.User
	err := r.DB.QueryRowContext(ctx, getUserByIdQuery, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// GetUserByEmail:
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

// CreateUser:
func (r *userRepository) CreateUser(ctx context.Context, user *dtos.CreateUserDto) (*entity.User, error) {
	args := []interface{}{user.Name, user.Email, user.Password}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var u entity.User

	err := r.DB.QueryRowContext(ctx, createUserQuery, args...).Scan(
		&u.ID,
		&u.CreatedAt,
		&u.Name,
		&u.Email,
		&u.Password,
	)
	// TODO: Manejo del error: unique key constraint
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UpdateUser:
func (r *userRepository) UpdateUser(ctx context.Context, id int64, user *dtos.UpdateUserDto) (*entity.User, error) {
	args := []interface{}{user.Name, user.Email, user.Password, id}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var u entity.User
	err := r.DB.QueryRowContext(ctx, updateUserQuery, args...).Scan(
		&u.ID,
		&u.CreatedAt,
		&u.Name,
		&u.Email,
		&u.Password,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// DeleteUser:
func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := r.DB.ExecContext(ctx, deleteUserQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// TODO: Agregar el error ErrRecordNotFound
	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}
