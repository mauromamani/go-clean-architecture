package user

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/ent"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
)

type UseCase interface {
	Get(ctx context.Context) ([]*ent.User, error)
	GetById(ctx context.Context, id int) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	Create(ctx context.Context, user *dtos.CreateUserDto) (*ent.User, error)
	Update(ctx context.Context, id int, user *dtos.UpdateUserDto) (*ent.User, error)
	Delete(ctx context.Context, id int) error
}
