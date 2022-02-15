package user

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/ent"
)

type UseCase interface {
	Get(ctx context.Context) ([]*ent.User, error)
	GetById(ctx context.Context, id int) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	Create(ctx context.Context, user *ent.User) (*ent.User, error)
	Update(ctx context.Context, user *ent.User) (*ent.User, error)
	Delete(ctx context.Context, id int) error
}
