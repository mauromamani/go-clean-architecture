package user

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
	"github.com/mauromamani/go-clean-architecture/internal/user/entity"
)

type Repository interface {
	GetUser(ctx context.Context) (*entity.User, error)
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *dtos.CreateUserDto) (*entity.User, error)
	UpdateUser(ctx context.Context, id int, user *dtos.UpdateUserDto) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}
